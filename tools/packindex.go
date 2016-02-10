package main

import (
	"fmt"
	"io"
	"encoding/binary"
	"encoding/hex"
)

// Only Version 2..

type PackIndex struct {
	hash string   
	crc string
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
	hashOffset := 0x408 + int64(indexAt) * 0x14
	crcOffset := 0x408 + int64(count) * 0x14 + int64(indexAt) * 4
	indexOffset := 0x408 + int64(count) * 0x18 + int64(indexAt) * 4

	_, err = in.Seek(hashOffset, 0)
	hashBuff := make([]byte, 20)
	_, err = in.Read(hashBuff)
	if err != nil {
		return PackIndex{}, nil
	}

	_, err = in.Seek(crcOffset, 0)
	crcBuff := make([]byte, 4)
	_, err = in.Read(crcBuff)
	if err != nil {
		return PackIndex{}, nil
	}

	_, err = in.Seek(indexOffset, 0)
	indexBuff := make([]byte, 4)
	_, err = in.Read(indexBuff)
	if err != nil {
		return PackIndex{}, nil
	}

	return PackIndex{
		hash: hex.EncodeToString(hashBuff),
		crc: hex.EncodeToString(crcBuff),
		offset: int(binary.BigEndian.Uint32(indexBuff)) }, nil
}

