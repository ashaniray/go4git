package main

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// Only Version 2..

type PackIndex struct {
	hash   string
	crc    string
	offset int
}

func (idx PackIndex) String() string {
	return fmt.Sprintf("Offset: %d, Hash: %s, crc: %s", idx.offset, idx.hash, idx.crc)
}

func getTotalCount(in io.ReadSeeker) (uint, error) {
	_, err := in.Seek(0x404, 0)
	if err != nil {
		return 0, err
	}
	buff := make([]byte, 4)
	_, err = in.Read(buff)
	if err != nil {
		return 0, err
	}
	count := binary.BigEndian.Uint32(buff)
	return uint(count), nil
}

func ReadPackIndexAt(indexAt int, in io.ReadSeeker) (PackIndex, error) {
	count, err := getTotalCount(in)
	if err != nil {
		return PackIndex{}, err
	}
	hashOffset := 0x408 + int64(indexAt)*0x14
	crcOffset := 0x408 + int64(count)*0x14 + int64(indexAt)*4
	indexOffset := 0x408 + int64(count)*0x18 + int64(indexAt)*4

	_, err = in.Seek(hashOffset, 0)
	hashBuff := make([]byte, 20)
	_, err = in.Read(hashBuff)
	if err != nil {
		return PackIndex{}, err
	}

	_, err = in.Seek(crcOffset, 0)
	crcBuff := make([]byte, 4)
	_, err = in.Read(crcBuff)
	if err != nil {
		return PackIndex{}, err
	}

	_, err = in.Seek(indexOffset, 0)
	indexBuff := make([]byte, 4)
	_, err = in.Read(indexBuff)
	if err != nil {
		return PackIndex{}, err
	}

	return PackIndex{
		hash:   hex.EncodeToString(hashBuff),
		crc:    hex.EncodeToString(crcBuff),
		offset: int(binary.BigEndian.Uint32(indexBuff))}, nil
}

func getFanoutValue(index int64, in io.ReadSeeker) (int, error) {
	if index < 0 {
		return 0, nil
	}
	_, err := in.Seek(8+index*4, 0)
	indexBuff := make([]byte, 4)
	_, err = in.Read(indexBuff)
	if err != nil {
		return 0, err
	}
	return int(binary.BigEndian.Uint32(indexBuff)), nil
}

func GetObjectForHash(hash string, in io.ReadSeeker) (PackIndex, error) {
	if len(hash) > 40 {
		return PackIndex{}, errors.New("Hash length is greater than 20")
	}
	idx, err := strconv.ParseInt(hash[:2], 16, 32)
	if err != nil {
		return PackIndex{}, err
	}
	end, err := getFanoutValue(idx, in)
	if err != nil {
		return PackIndex{}, err
	}
	start, err := getFanoutValue(idx-1, in)
	if err != nil {
		return PackIndex{}, err
	}

	for start <= end {
		curr := (start + end) / 2
		indexObj, err := ReadPackIndexAt(curr, in)
		if err != nil {
			return PackIndex{}, err
		}
		//switch (strings.Compare(hash, indexObj.hash[:len(hash)])) {
		sliceHash := indexObj.hash[:len(hash)]
		if hash == sliceHash {
			return indexObj, nil
		}
		if hash > sliceHash {
			if start < curr {
				start = curr
			} else {
				start++
			}
		} else {
			if end > curr {
				end = curr
			} else {
				end--
			}
		}
	}
	return PackIndex{}, errors.New("Object not found in index file")
}
