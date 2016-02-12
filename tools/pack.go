package main

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
	objectType  ObjectType
	data        []byte
	hashOfRef   string
	refOffset   int64
	size        int64
	endOffset   int64
	startOffset int64
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
	str := fmt.Sprintf("%s %s %d %d %d", o.getHash(), o.objectType, o.size, o.endOffset-o.startOffset, o.startOffset)
	str += fmt.Sprintf("  Offset end: %d", o.endOffset)
	return str
}

func (o PackedObject) getHash() string {
	b, _ := GenSHA1(bytes.NewReader(o.data), o.objectType.String())
	return hex.EncodeToString(b)
}

func ReadPackedObjectAtOffset(offset int64, in io.ReadSeeker) (PackedObject, error) {
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
	objectType := ObjectType((int(headByte[0]) & 0x70) >> 4)
	objectSize := (int64(headByte[0])) & int64(0x0f)
	size, err := readVariableSize(in)
	if err != nil {
		return PackedObject{}, err
	}
	objectSize = objectSize + (size << 4)

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

	currOffset, _ := in.Seek(0, os.SEEK_CUR)

	return PackedObject{objectType: objectType,
		size:        objectSize,
		refOffset:   offset - negOffset,
		hashOfRef:   hashOfRef,
		data:        buff,
		startOffset: offset,
		endOffset:   currOffset,
	}, err
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

func readPackedBasicObjectData(in io.Reader, objectSize int64) ([]byte, error) {
	buff := make([]byte, objectSize, objectSize)
	zr, err := zlib.NewReader(in)
	if err != nil {
		return nil, err
	}
	defer zr.Close()
	_, err = zr.Read(buff)
	if err != nil && err != io.EOF {
		return nil, err
	}
	return buff, nil
}

func readRefDeltaObjectData(in io.Reader, objectSize int64) (string, error) {
	hashObject := make([]byte, 20)
	n, err := in.Read(hashObject)
	if err != nil {
		return "", err
	}
	return string(hashObject[:n]), err
}

func readOfsDeltaObjectData(in io.Reader, objectSize int64) (int64, error) {
	negativeOffset, err := readVariableSizeForOFS(in)
	if err != nil {
		return 0, err
	}
	return negativeOffset, err
}
