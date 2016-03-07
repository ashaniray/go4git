package go4git

import (
	"io"
	"bytes"
	"fmt"
)

type IndexEntry struct {
	CtimeSecs   uint32
	CtimeNanoSecs uint32
	MtimeSecs uint32
	MtimeNanoSecs uint32
	Dev uint32
	Ino uint32
	ObjectType int
	UnixPermission int
	Uid uint32
	Gid uint32
	Hash []byte
	Flags []byte
	AdditionalFlag []byte
	EntryPathName string
	Padding []byte
}

type IndexHeader struct {
	Signature string
	Version uint32
	CountEntries int
}

type Index struct {
	Header IndexHeader
	Entries []IndexEntry
}

func ParseIndex(in io.ReadSeeker) (Index, error) {
	header, err := parseIndexHeader(in)
	if err != nil {
		return Index{}, err
	}

	entries := make([]IndexEntry, header.CountEntries)
	for i := 0; i < header.CountEntries; i++ {
		entry, err := parseIndexEntry(in)
		if err != nil {
			return Index{}, err
		}
		entries[i] = entry
	}
	return Index{Header: header, Entries: entries}, nil
}

func parseIndexEntry(in io.Reader) (IndexEntry, error) {
	return IndexEntry {
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
		Signature:signature, 
		Version:version, 
		CountEntries:int(count),
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
	f := "CtimeSecs:%d\nCtimeNanoSecs:%d\nMtimeSecs:%d\nMtimeNanoSecs:%d\n" 
	f += "Dev:%d\nIno:%d\nObjectType:%d\nUnixPermission:%d\nUid:%d\n"
	f += "Gid:%d\nHash:%s\nFlags:%s\nAdditionalFlag:%s\nEntryPathName:%s\n"
	f += "Padding:%s\n"

	str := fmt.Sprintf(f,
		e.CtimeSecs,
		e.CtimeNanoSecs,
		e.MtimeSecs,
		e.MtimeNanoSecs,
		e.Dev,
		e.Ino,
		e.ObjectType,
		e.UnixPermission,
		e.Uid,
		e.Gid,
		Byte2String(e.Hash),
		Byte2String(e.Flags),
		Byte2String(e.AdditionalFlag),
		e.EntryPathName,
		Byte2String(e.Padding),
		)

	return str
}

func (idx Index) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("%s", idx.Header))
	for i, idx := range idx.Entries {
		buffer.WriteString(fmt.Sprintf("\nEntry [%d]:\n%s", i, idx))
	}
	return buffer.String()
}

