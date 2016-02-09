package main

import (
	"io"
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

func ReadPackedDataAtOffset(offset int64, in io.ReadSeeker) (ObjectType, int, []byte, error) {
	_, err := in.Seek(offset, 0)
	if err != nil {
		return 0, 0, nil, err
	}
	headByte := make([]byte, 1, 1)
	_, err = in.Read(headByte)
	if err != nil {
		return 0, 0, nil, err
	}


	objectType := (int(headByte[0]) & 0x70) >> 4

	size := (int(headByte[0])) & int(0x0f)

	var shiftBit uint = 4

	for {
		sizeByte := make([]byte, 1, 1)
		_, err = in.Read(sizeByte)
		if err != nil {
			return 0, 0, nil, err
		}
		sizeByteInInt := int(sizeByte[0])
		size = size + (sizeByteInInt << shiftBit)
		shiftBit += 8
		cont := (sizeByteInInt & 0x80) >> 7
		if cont == 0 {
			break
		}
	}

	buff := make([]byte, size, size)
	_, err = in.Read(buff)
	if err != nil {
		return 0, 0, nil, err
	}

	return ObjectType(objectType), size, buff, err
}

