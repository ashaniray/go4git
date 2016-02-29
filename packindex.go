package go4git

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Only Version 2..

type PackIndex struct {
	Hash   []byte
	CRC    []byte
	Offset int
}

type ByOffset []PackIndex

func (a ByOffset) Len() int           { return len(a) }
func (a ByOffset) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByOffset) Less(i, j int) bool { return a[i].Offset < a[j].Offset }

func (idx PackIndex) String() string {
	return fmt.Sprintf("%d %s (%s)", idx.Offset, Byte2String(idx.Hash), Byte2String(idx.CRC))
}

func (idx PackIndex) CRCAsString() string {
	return Byte2String(idx.CRC)
}

func (idx PackIndex) HashAsString() string {
	return Byte2String(idx.Hash)
}

func GetTotalCount(in io.ReadSeeker) (uint, error) {
	_, err := in.Seek(0x404, os.SEEK_SET)
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

func readHashAt(indexAt int, in io.ReadSeeker) ([]byte, error) {
	hashOffset := 0x408 + int64(indexAt)*0x14
	_, err := in.Seek(hashOffset, os.SEEK_SET)
	if err != nil {
		return nil, err
	}
	hashBuff := make([]byte, 20)
	_, err = in.Read(hashBuff)
	if err != nil {
		return nil, err
	}
	return hashBuff, nil
}

func readCRCAt(indexAt int, count uint, in io.ReadSeeker) ([]byte, error) {
	crcOffset := 0x408 + int64(count)*0x14 + int64(indexAt)*4
	_, err := in.Seek(crcOffset, os.SEEK_SET)
	if err != nil {
		return nil, err
	}
	crcBuff := make([]byte, 4)
	_, err = in.Read(crcBuff)
	if err != nil {
		return nil, err
	}
	return crcBuff, nil
}

func readOffsetAt(indexAt int, count uint, in io.ReadSeeker) (int, error) {
	indexOffset := 0x408 + int64(count)*0x18 + int64(indexAt)*4
	_, err := in.Seek(indexOffset, os.SEEK_SET)
	if err != nil {
		return 0, err
	}
	indexBuff := make([]byte, 4)
	_, err = in.Read(indexBuff)
	if err != nil {
		return 0, err
	}
	return int(binary.BigEndian.Uint32(indexBuff)), nil
}

func GetAllPackedIndex(in io.ReadSeeker) ([]PackIndex, error) {
	count, err := GetTotalCount(in)
	if err != nil {
		panic(err)
	}
	indices := make([]PackIndex, count)
	for i := 0; i < int(count); i++ {
		// TODO: Optimize the reading.. The reading of ALL indices
		// can be done serially. Instead of calling ReadPackIndexAt
		// for each i, ALL the data can be read directly.
		indices[i], err = ReadPackIndexAt(i, in)
		if err != nil {
			panic(err)
		}
	}
	return indices, err
}

func ReadPackIndexAt(indexAt int, in io.ReadSeeker) (PackIndex, error) {
	count, err := GetTotalCount(in)
	if err != nil {
		return PackIndex{}, err
	}

	hash, err := readHashAt(indexAt, in)
	if err != nil {
		return PackIndex{}, err
	}

	crc, err := readCRCAt(indexAt, count, in)
	if err != nil {
		return PackIndex{}, err
	}

	offset, err := readOffsetAt(indexAt, count, in)
	if err != nil {
		return PackIndex{}, err
	}

	return PackIndex{
		Hash:   hash,
		CRC:    crc,
		Offset: offset,
	}, nil
}

func getFanoutValue(index int64, in io.ReadSeeker) (int, error) {
	if index < 0 {
		return 0, nil
	}
	_, err := in.Seek(8+index*4, os.SEEK_SET)
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
		hashOfObj, err := readHashAt(curr, in)
		if err != nil {
			return PackIndex{}, err
		}
		hashOfObjAsStr := Byte2String(hashOfObj)
		switch strings.Compare(hash, hashOfObjAsStr[:len(hash)]) {
		case 0:
			return ReadPackIndexAt(curr, in)
		case 1:
			if start < curr {
				start = curr
			} else {
				start++
			}
		case -1:
			if end > curr {
				end = curr
			} else {
				end--
			}
		}
	}
	return PackIndex{}, errors.New("Object not found in index file")
}
