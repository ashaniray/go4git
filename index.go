package go4git

import (
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
	signature string
	version uint32
	countEntries uint32
}

type Index struct {
	header IndexHeader
	entries []IndexEntry
}


func (idx IndexEntry) String() string {
	return "TODO"
}

