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
	size           int
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
	headByte := make([]byte, 1)
	_, err := in.Read(headByte)
	if err != nil {
		return PackedObject{}, err
	}

	objectType := ObjectType((int(headByte[0]) & 0x70) >> 4)
	objectSize := (int(headByte[0])) & int(0x0f)
	var shiftBit uint = 4
	for {
		sizeByte := make([]byte, 1)
		_, err = in.Read(sizeByte)
		if err != nil {
			return PackedObject{}, err
		}
		sizeByteInInt := int(sizeByte[0])
		objectSize = objectSize + ((sizeByteInInt & 0x7f) << shiftBit)
		shiftBit += 7
		cont := (sizeByteInInt & 0x80) >> 7
		if cont == 0 {
			break
		}
	}

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

func readPackedBasicObjectData(in io.Reader, objectSize int) ([]byte, error) {
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

func readRefDeltaObjectData(in io.Reader, objectSize int) ([]byte, string, error) {
	hashObject := make([]byte, 20)
	n, err := in.Read(hashObject)
	if err != nil {
		return nil, "", err
	}
	fmt.Printf("Read Hash: %v\n", hashObject)

	buff, err := readPackedBasicObjectData(in, objectSize)
	return buff, string(hashObject[:n]), err
}

//From documentation:
//n bytes with MSB set in all but the last one.
//The offset is then the number constructed by
//concatenating the lower 7 bit of each byte, and
//for n >= 2 adding 2^7 + 2^14 + ... + 2^(7*(n-1))
//to the result.
func readOfsDeltaObjectData(in io.Reader, objectSize int) ([]byte, int64, error) {
	var negativeOffset int64 = 0
	var shiftBit uint = 0
	for {
		sizeByte := make([]byte, 1)
		_, err := in.Read(sizeByte)
		if err != nil {
			return nil, 0, err
		}
		sizeByteInInt := int64(sizeByte[0])
		negativeOffset = negativeOffset + ((sizeByteInInt & 0x7f) << shiftBit)
		shiftBit += 7
		cont := (sizeByteInInt & 0x80) >> 7
		if cont == 0 {
			break
		}
	}
	buff, err := readPackedBasicObjectData(in, objectSize)
	return buff, negativeOffset, err
}
