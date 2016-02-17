package go4git

import (
	"bytes"
	"compress/zlib"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

type ObjectType int

type PackedObject struct {
	Type        ObjectType
	Data        []byte
	HashOfRef   string
	RefOffset   int64
	Size        int64
	StartOffset int64
	DeltaData   []byte
}

const (
	_ = iota
	COMMIT
	TREE
	BLOB
	TAG
	_
	OFS_DELTA
	REF_DELTA
)

func (t ObjectType) String() string {
	switch t {
	case COMMIT:
		return "commit"
	case TREE:
		return "tree"
	case BLOB:
		return "blob"
	case TAG:
		return "tag"
	case REF_DELTA:
		return "ref_delta"
	case OFS_DELTA:
		return "ofs_delta"
	}
	return "Unknown type"
}

func (o PackedObject) String() string {
	str := fmt.Sprintf("%s %s %d %d", o.GetHash(), o.Type, o.Size, o.StartOffset)
	return str
}

func (o PackedObject) GetHash() string {
	b, _ := GenSHA1(bytes.NewReader(o.Data), o.Type.String())
	return hex.EncodeToString(b)
}

func ReadPackedObjectAtOffset(offset int64, in io.ReadSeeker, inIndex io.ReadSeeker) (PackedObject, error) {
	_, err := in.Seek(offset, os.SEEK_SET)
	if err != nil {
		return PackedObject{}, err
	}

	//Read the header and size
	headByte := make([]byte, 1)
	_, err = in.Read(headByte)
	if err != nil {
		return PackedObject{}, err
	}

	firstByte := int(headByte[0])
	objectType := ObjectType((firstByte & 0x70) >> 4)
	objectSize := (int64(firstByte)) & int64(0x0f)

	if firstByte & 0x80 > 0 {
		size, err := readVariableSize(in)
		if err != nil {
			return PackedObject{}, err
		}
		objectSize = objectSize + (size << 4)
	}

	// Read the data
	var hashOfRef string
	var negOffset int64
	switch objectType {
	case REF_DELTA:
		hashOfRef, err = readRefDeltaObjectData(in, objectSize)
	case OFS_DELTA:
		negOffset, err = readOfsDeltaObjectData(in, objectSize)
	}
	buff, err := readPackedBasicObjectData(in, objectSize)

	obj := PackedObject{
		Type:        objectType,
		Size:        objectSize,
		RefOffset:   offset - negOffset,
		HashOfRef:   hashOfRef,
		Data:        buff,
		StartOffset: offset,
		DeltaData:   nil,
	}
	// Patch the deltas....
	if objectType == OFS_DELTA {
		base, err := ReadPackedObjectAtOffset(offset-negOffset, in, inIndex)
		if err != nil {
			return obj, err
		}
		targetBuff := applyDeltaBuffer(base.Data, buff)
		obj.DeltaData = buff
		obj.Data = targetBuff
	}
	if objectType == REF_DELTA {
		packedIndex, err := GetObjectForHash(hashOfRef, inIndex)
		if err == nil {
			return obj, err
		}
		base, err := ReadPackedObjectAtOffset(int64(packedIndex.Offset), in, inIndex)
		if err != nil {
			return obj, err
		}
		targetBuff := applyDeltaBuffer(base.Data, buff)
		obj.DeltaData = buff
		obj.Data = targetBuff
	}
	return obj, err

}

func readVariableSize(in io.Reader) (int64, error) {
	var size int64 = 0
	var shiftBit uint = 0
	for {
		sizeByte := make([]byte, 1)
		_, err := in.Read(sizeByte)
		if err != nil {
			return 0, err
		}
		sizeByteInInt := int64(sizeByte[0])
		size = size + ((sizeByteInInt & 0x7f) << shiftBit)
		shiftBit += 7
		cont := (sizeByteInInt & 0x80) >> 7
		if cont == 0 {
			break
		}
	}
	return size, nil
}

func readVariableSizeForOFS(in io.Reader) (int64, error) {
	sizeByte := make([]byte, 1)
	_, err := in.Read(sizeByte)
	if err != nil {
		return 0, err
	}
	c := int64(sizeByte[0])
	size := c & 0x7f
	for (c & 0x80) != 0 {
		size += 1
		_, err = in.Read(sizeByte)
		if err != nil {
			return 0, err
		}
		c = int64(sizeByte[0])
		size = (size << 7) + (c & 0x7f)
	}
	return size, nil
}

func readPackedBasicObjectData(in io.ReadSeeker, objectSize int64) ([]byte, error) {
	buff := make([]byte, objectSize)
	zr, err := zlib.NewReader(in)
	if err != nil {
		return nil, err
	}
	defer zr.Close()
	n, err := zr.Read(buff)
	if err != nil && err != io.EOF {
		return nil, err
	}
	return buff[:n], nil
}

func readRefDeltaObjectData(in io.Reader, objectSize int64) (string, error) {
	hashObject := make([]byte, 20)
	n, err := in.Read(hashObject)
	if err != nil {
		return "", err
	}
	return string(hashObject[:n]), err
}

func readSizeInDelta(delta []byte) (size int, bytesRead uint) {
	if len(delta) == 0 {
		return 0, 0
	}
	for {
		c := int(delta[bytesRead])
		tmpSize := c & 0x7f
		size = size | (tmpSize << (7 * bytesRead))
		bytesRead++
		if (c & 0x80) == 0 {
			break
		}
	}
	return size, bytesRead
}

func getCopyOrPasteInDelta(delta []byte) (isCopy bool, srcOffset uint, length uint, bytesRead uint) {
	if len(delta) == 0 {
		return
	}
	c := delta[bytesRead]
	bytesRead++
	switch (c & 0x80) >> 7 {
	case 1:
		isCopy = true
		for i := uint(0); i < 4; i++ {
			if c&(1<<i) != 0 {
				b := delta[bytesRead]
				bytesRead++
				srcOffset = (uint(b) << uint(8*i)) | srcOffset
			}
		}
		for i := uint(0); i < 3; i++ {
			if c&(1<<(i+4)) != 0 {
				b := delta[bytesRead]
				bytesRead++
				length = (uint(b) << uint(8*i)) | length
			}
		}
	case 0:
		isCopy = false
		length = uint(delta[0])
	default:
		panic("Impossible code")
	}
	return
}

func applyDeltaBuffer(src []byte, delta []byte) []byte {
	if len(delta) == 0 {
		return src
	}
	var deltaOffset uint = 0
	srcSize, bytesRead := readSizeInDelta(delta[deltaOffset:])
	deltaOffset += bytesRead

	destSize, bytesRead := readSizeInDelta(delta[deltaOffset:])
	deltaOffset += bytesRead

	_, _ = srcSize, destSize

	target := make([]byte, destSize)
	var targetOffset uint = 0

	for {
		isCopy, srcOffset, length, bytesRead := getCopyOrPasteInDelta(delta[deltaOffset:])
		deltaOffset += bytesRead
		if length < 1 {
			break
		}
		var buff []byte
		if isCopy {
			buff = src[srcOffset : srcOffset+length]
		} else {
			buff = delta[deltaOffset : deltaOffset+length]
			deltaOffset += length
		}
		copy(target[targetOffset:], buff)
		targetOffset += length
		if len(delta) <= int(deltaOffset) {
			break
		}
	}
	return target
}

func readOfsDeltaObjectData(in io.Reader, objectSize int64) (int64, error) {
	negativeOffset, err := readVariableSizeForOFS(in)
	if err != nil {
		return 0, err
	}
	return negativeOffset, err
}

