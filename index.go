package go4git

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

type IndexEntry struct {
	CtimeSecs      uint32
	CtimeNanoSecs  uint32
	MtimeSecs      uint32
	MtimeNanoSecs  uint32
	Dev            uint32
	Ino            uint32
	ObjectType     int
	UnixPermission int
	Uid            uint32
	Gid            uint32
	FileSize       int
	Hash           []byte
	Flags          []byte
	AdditionalFlag []byte
	EntryPathName  string
	Padding        []byte
}

type IndexHeader struct {
	Signature    string
	Version      uint32
	CountEntries int
}

type Index struct {
	Header  IndexHeader
	Entries []IndexEntry
}

func ParseIndex(in io.ReadSeeker) (Index, error) {
	header, err := parseIndexHeader(in)
	if err != nil {
		return Index{}, err
	}

	entries := make([]IndexEntry, header.CountEntries)
	inIndex := bufio.NewReader(in)
	for i := 0; i < header.CountEntries; i++ {
		entry, err := parseIndexEntry(inIndex)
		if err != nil {
			return Index{}, err
		}
		entries[i] = entry
	}
	return Index{Header: header, Entries: entries}, nil
}

func parseIndexEntry(in *bufio.Reader) (IndexEntry, error) {

	bytesRead := 0
	ctimeSecs, err := ReadUint32(in)
	if err != nil {
		return IndexEntry{}, err
	}
	bytesRead += 4
	ctimeNanoSecs, err := ReadUint32(in)
	if err != nil {
		return IndexEntry{}, err
	}
	bytesRead += 4
	mtimeSecs, err := ReadUint32(in)
	if err != nil {
		return IndexEntry{}, err
	}
	bytesRead += 4
	mtimeNanoSecs, err := ReadUint32(in)
	if err != nil {
		return IndexEntry{}, err
	}
	bytesRead += 4
	dev, err := ReadUint32(in)
	if err != nil {
		return IndexEntry{}, err
	}
	bytesRead += 4
	ino, err := ReadUint32(in)
	if err != nil {
		return IndexEntry{}, err
	}
	bytesRead += 4
	mode, err := ReadUint32(in)
	if err != nil {
		return IndexEntry{}, err
	}
	bytesRead += 4
	objectType := (mode & 0xf000) >> 12
	unixPermission := mode & 0x01ff
	uid, err := ReadUint32(in)
	if err != nil {
		return IndexEntry{}, err
	}
	bytesRead += 4
	gid, err := ReadUint32(in)
	if err != nil {
		return IndexEntry{}, err
	}
	bytesRead += 4
	fileSize, err := ReadUint32(in)
	if err != nil {
		return IndexEntry{}, err
	}
	bytesRead += 4
	sha1 := make([]byte, 20)
	_, err = in.Read(sha1)
	if err != nil {
		return IndexEntry{}, err
	}
	bytesRead += 20
	flags := make([]byte, 2)
	_, err = in.Read(flags)
	if err != nil {
		return IndexEntry{}, err
	}
	bytesRead += 2
	// TODO: Version 3 has another flag

	entryPathName, err := in.ReadString(0)
	if err != nil {
		return IndexEntry{}, err
	}
	entryPathName = entryPathName[0 : len(entryPathName)-1]
	bytesRead += len(entryPathName) + 1
	padSize := 8 - (bytesRead % 8)
	if padSize == 8 {
		padSize = 0
	}

	padding := make([]byte, padSize)
	_, err = in.Read(padding)
	if err != nil {
		return IndexEntry{}, err
	}
	bytesRead += len(entryPathName) + 1

	return IndexEntry{
		CtimeSecs:      ctimeSecs,
		CtimeNanoSecs:  ctimeNanoSecs,
		MtimeSecs:      mtimeSecs,
		MtimeNanoSecs:  mtimeNanoSecs,
		Dev:            dev,
		Ino:            ino,
		ObjectType:     int(objectType),
		UnixPermission: int(unixPermission),
		Uid:            uid,
		Gid:            gid,
		FileSize:       int(fileSize),
		Hash:           sha1,
		Flags:          flags,
		AdditionalFlag: nil,
		EntryPathName:  entryPathName,
		Padding:        padding,
	}, nil
}

func parseIndexHeader(in io.Reader) (IndexHeader, error) {
	buff := make([]byte, 4)
	n, err := in.Read(buff)
	if err != nil {
		return IndexHeader{}, err
	}
	signature := string(buff[:n])

	version, err := ReadUint32(in)
	if err != nil {
		return IndexHeader{}, err
	}

	count, err := ReadUint32(in)
	if err != nil {
		return IndexHeader{}, err
	}

	return IndexHeader{
		Signature:    signature,
		Version:      version,
		CountEntries: int(count),
	}, nil
}

func (hdr IndexHeader) String() string {
	str := fmt.Sprintf("Signature:%s Version:%d Entries:%d",
		hdr.Signature,
		hdr.Version,
		hdr.CountEntries)
	return str
}

func (e IndexEntry) String() string {
	f := "%o%04o %s %d %s"
	str := fmt.Sprintf(f,
		e.ObjectType,
		e.UnixPermission,
		Byte2String(e.Hash),
		0,
		e.EntryPathName,
	)
	return str
}

func (idx Index) String() string {
	var buffer bytes.Buffer
	for _, idx := range idx.Entries {
		buffer.WriteString(fmt.Sprintf("%s\n", idx))
	}
	return buffer.String()
}
