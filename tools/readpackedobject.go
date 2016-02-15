package main

import (
	"flag"
	"fmt"
	"os"
	"io"
	"sort"
)

var offset = flag.Int64("s", -1, "The offset to read from the pack file")
var verbose = flag.Bool("t", false, "Output verbose information")
var verify = flag.Bool("v", true, "Produce output of git pack-verify -v")

func showVerifyPack(inPack io.ReadSeeker, inIdx io.ReadSeeker) {
	indices, err := GetAllPackedIndex(inIdx)
	if err != nil {
		panic(err)
	}
	sort.Sort(ByOffset(indices))
	cnt := len(indices)
	o, err := ReadPackedObjectAtOffset(int64(indices[0].Offset), inPack)
	if err != nil {
		panic(err)
	}
	for i := 0; i < cnt - 1; i++ {
		next, _ := ReadPackedObjectAtOffset(int64(indices[i + 1].Offset), inPack)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s %s %d %d %d\n", 
			o.GetHash(), 
			o.objectType, 
			o.size,
			next.startOffset - o.startOffset,
			o.startOffset)
		o = next
	}
}

func main() {
	flag.Parse()
	f, err := GetArgInputFile()
	if err != nil {
		panic(err)
	}
	packFile := f.Name()
	idxName := packFile[:len(packFile) - 4] + "idx"
	inIdx, err := os.Open(idxName)
	if err != nil {
		panic(err)
	}
	showVerifyPack(f, inIdx)
	return
	p, err := ReadPackedObjectAtOffset(*offset, f)
	if err != nil {
		panic(err)
	}
	if *verbose {
		fmt.Fprintf(os.Stdout, "Object at [%d] => Type: %s, Size: %d\n", *offset, p.objectType, p.size)
		fmt.Fprintf(os.Stdout, "  ObjRef: %s, ObjOffset: %d\n", p.hashOfRef, p.refOffset)
		fmt.Fprintf(os.Stdout, "  Data(starts below):\n")
	}
	fmt.Fprintf(os.Stdout, "%s", p.data)
}
