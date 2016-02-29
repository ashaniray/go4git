package go4git

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type ObjectType int

type PackedObject struct {
	Type        ObjectType
	Data        []byte
	HashOfRef   []byte
	RefOffset   int64
	Size        int64
	StartOffset int64
	DeltaData   []byte
	ActualType  ObjectType
	Hash        []byte
	RefLevel    int
	BaseHash    []byte
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
	str := fmt.Sprintf("Packed Object at offset [%d] \n", o.StartOffset)
	str += fmt.Sprintf(" Type: %s\n", o.Type)
	str += fmt.Sprintf(" HashOfRef: %s\n", Byte2String(o.HashOfRef))
	str += fmt.Sprintf(" RefOffset: %d\n", o.RefOffset)
	str += fmt.Sprintf(" Size: %d\n", o.Size)
	str += fmt.Sprintf(" StartOffset: %d\n", o.StartOffset)
	str += fmt.Sprintf(" ActualType: %s\n", o.ActualType)
	str += fmt.Sprintf(" Hash: %s\n", Byte2String(o.Hash))
	str += fmt.Sprintf(" RefLevel: %d\n", o.RefLevel)
	str += fmt.Sprintf(" BaseHash: %s\n", Byte2String(o.BaseHash))
	str += fmt.Sprintf(" ---Data(starts below):---\n")
	str += fmt.Sprintf("%s", o.Data)
	return str
}

func (o PackedObject) getHash() []byte {
	b, _ := GenSHA1(bytes.NewReader(o.Data), o.ActualType.String())
	return b
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

	if firstByte&0x80 > 0 {
		size, err := readVariableSize(in)
		if err != nil {
			return PackedObject{}, err
		}
		objectSize = objectSize + (size << 4)
	}

	// Read the data
	var hashOfRef []byte
	var negOffset int64
	switch objectType {
	case REF_DELTA:
		hashOfRef, err = readRefDeltaObjectData(in, objectSize)
	case OFS_DELTA:
		negOffset, err = readOfsDeltaObjectData(in, objectSize)
	}

	buff, err := UnzlibToBuffer(in)

	obj := PackedObject{
		Type:        objectType,
		Size:        objectSize,
		RefOffset:   offset - negOffset,
		HashOfRef:   hashOfRef,
		Data:        buff,
		StartOffset: offset,
		DeltaData:   nil,
		RefLevel:    0,
		ActualType:  objectType,
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
		obj.ActualType = base.ActualType
		obj.RefLevel = base.RefLevel + 1
		obj.BaseHash = base.Hash
	}

	if objectType == REF_DELTA {
		packedIndex, err := GetObjectForHash(Byte2String(hashOfRef), inIndex)
		if err != nil {
			return obj, err
		}
		base, err := ReadPackedObjectAtOffset(int64(packedIndex.Offset), in, inIndex)
		if err != nil {
			return obj, err
		}
		targetBuff := applyDeltaBuffer(base.Data, buff)
		obj.DeltaData = buff
		obj.Data = targetBuff
		obj.ActualType = base.ActualType
		obj.RefLevel = base.RefLevel + 1
		obj.BaseHash = base.Hash
	}
	obj.Hash = obj.getHash()
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

func readRefDeltaObjectData(in io.Reader, objectSize int64) ([]byte, error) {
	hashObject := make([]byte, 20)
	n, err := in.Read(hashObject)
	if err != nil {
		return nil, err
	}
	return hashObject[:n], err
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
		if length == 0 {
			length = 0x010000
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
