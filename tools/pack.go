package main

import (
	"compress/zlib"
	"fmt"
	"io"
)

type ObjectType int

type PackedObject struct {
	objectType     ObjectType
	data           []byte
	hashOfRef      string
	negOffsetOfRef int64
	size           int64
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

func ReadPackedObjectAtOffset(offset int64, in io.ReadSeeker) (PackedObject, error) {
	_, err := in.Seek(offset, 0)
	if err != nil {
		return PackedObject{}, err
	}
	return ReadPackedObject(in)
}

func ReadPackedObject(in io.ReadSeeker) (PackedObject, error) {
	//Read the header and size
	headByte := make([]byte, 1)
	_, err := in.Read(headByte)
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
	var buff []byte
	var hashOfRef string
	var negOffset int64
	switch objectType {
	case TREE, BLOB, COMMIT:
		buff, err = readPackedBasicObjectData(in, objectSize)
	case REF_DELTA:
		buff, hashOfRef, err = readRefDeltaObjectData(in, objectSize)
	case OFS_DELTA:
		buff, negOffset, err = readOfsDeltaObjectData(in, objectSize)
	}

	return PackedObject{objectType: objectType, size: objectSize, negOffsetOfRef: negOffset, hashOfRef: hashOfRef, data: buff}, err
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

func readPackedBasicObjectData(in io.Reader, objectSize int64) ([]byte, error) {
	buff := make([]byte, objectSize)
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

func readRefDeltaObjectData(in io.Reader, objectSize int64) ([]byte, string, error) {
	hashObject := make([]byte, 20)
	n, err := in.Read(hashObject)
	if err != nil {
		return nil, "", err
	}
	fmt.Printf("Read Hash: %v\n", hashObject)

	buff, err := readPackedBasicObjectData(in, objectSize)
	return buff, string(hashObject[:n]), err
}

func readOfsDeltaObjectData(in io.Reader, objectSize int64) ([]byte, int64, error) {
	negativeOffset, err := readVariableSize(in)
	if err != nil {
		return nil, 0, err
	}
	buff, err := readPackedBasicObjectData(in, objectSize)
	return buff, negativeOffset, err
}
