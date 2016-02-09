package main

import (
	"io"
	"compress/zlib"
)


type ObjectType int

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
	case DELTA1:
		return "delta1"
	case DELTA2:
		return "delta2"
	}
	return "Unknown type"
}

const (
	_ = iota
	COMMIT
	TREE
	BLOB
	TAG
	_
	DELTA1
	DELTA2
)


func ReadPackedObjectAtOffset(offset int64, in io.ReadSeeker) (ObjectType, int, []byte, error) {
	_, err := in.Seek(offset, 0)
	if err != nil {
		return 0, 0, nil, err
	}
	return ReadPackedObject(in)
}

func ReadPackedObject(in io.ReadSeeker) (ObjectType, int, []byte, error) {
	headByte := make([]byte, 1, 1)
	_, err := in.Read(headByte)
	if err != nil {
		return 0, 0, nil, err
	}


	objectType := ObjectType((int(headByte[0]) & 0x70) >> 4)
	objectSize := (int(headByte[0])) & int(0x0f)
	var shiftBit uint = 4
	for {
		sizeByte := make([]byte, 1, 1)
		_, err = in.Read(sizeByte)
		if err != nil {
			return 0, 0, nil, err
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
	switch (objectType) {
		case TREE, BLOB, COMMIT:
			buff, err = readPackedBasicObjectData(in, objectSize)
		case DELTA1:
		case DELTA2:
	}

	return objectType, objectSize, buff, err
}

func readPackedBasicObjectData(in io.Reader, objectSize int) ([]byte, error) {
	buff := make([]byte, objectSize)
	zr, err := zlib.NewReader(in)
	if err != nil {
		return nil, err
	}
	defer zr.Close()
	n, err := zr.Read(buff)
	if err != nil {
		if err == io.EOF {
			err = nil
		} else {
			return nil, err
		}
	}
	buff = buff[:n]
	return buff, nil
}

