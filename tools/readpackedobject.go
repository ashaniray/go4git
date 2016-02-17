package main

import (
	"flag"
	"fmt"
	"github.com/ashaniray/go4git"
	"io"
	"os"
	"sort"
)

var offset = flag.Int64("s", -1, "The offset to read from the pack file")
var verbose = flag.Bool("t", false, "Output verbose information")
var verifyPack = flag.Bool("v", true, "Produce output of git pack-verify -v")
var debug = flag.Bool("d", false, "Outputs debug information to debug the code")

func showDebugInfo(inPack io.ReadSeeker, inIdx io.ReadSeeker) {
	indices, err := go4git.GetAllPackedIndex(inIdx)
	if err != nil {
		panic(err)
	}
	sort.Sort(go4git.ByOffset(indices))
	cnt := len(indices)
	o, err := go4git.ReadPackedObjectAtOffset(int64(indices[0].Offset), inPack, inIdx)
	if err != nil {
		panic(err)
	}
	for i := 0; i < cnt-1; i++ {
		next, _ := go4git.ReadPackedObjectAtOffset(int64(indices[i+1].Offset), inPack, inIdx)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(os.Stdout, "################ Data at %d #############\n", i)
		fmt.Fprintf(os.Stdout, "%s %s %d %d %d %d Data(Below)\n%s\n",
			o.Hash,
			o.Type,
			o.Size,
			next.StartOffset-o.StartOffset,
			o.StartOffset,
			o.RefOffset,
			o.Data,
		)
		o = next
	}
}

func showVerifyPack(inPack io.ReadSeeker, inIdx io.ReadSeeker) {
	indices, err := go4git.GetAllPackedIndex(inIdx)
	if err != nil {
		panic(err)
	}
	sort.Sort(go4git.ByOffset(indices))
	cnt := len(indices)
	o, err := go4git.ReadPackedObjectAtOffset(int64(indices[0].Offset), inPack, inIdx)
	if err != nil {
		panic(err)
	}
	for i := 0; i < cnt-1; i++ {
		next, _ := go4git.ReadPackedObjectAtOffset(int64(indices[i+1].Offset), inPack, inIdx)
		if err != nil {
			panic(err)
		}
		str := fmt.Sprintf("%s %s\t%d %d %d", o.Hash, o.ActualType, o.Size, next.StartOffset - o.StartOffset, o.StartOffset)
		if o.Type == go4git.REF_DELTA || o.Type == go4git.OFS_DELTA {
			str += fmt.Sprintf(" %d %s", o.RefLevel, o.BaseHash)
		}
		fmt.Fprintf(os.Stdout, "%s\n", str)
		o = next
	}
}

func main() {
	flag.Parse()
	f, err := go4git.GetArgInputFile()
	if err != nil {
		panic(err)
	}
	packFile := f.Name()
	idxName := packFile[:len(packFile)-4] + "idx"
	inIdx, err := os.Open(idxName)
	if err != nil {
		panic(err)
	}

	// TODO: This in only for easy debugging purposes..
	// To be deleted...
	if *debug {
		showDebugInfo(f, inIdx)
		return
	}

	if *offset != -1 {
		p, err := go4git.ReadPackedObjectAtOffset(*offset, f, inIdx)
		if err != nil {
			panic(err)
		}
		if *verbose {
			fmt.Fprintf(os.Stdout, "Object at [%d] => Type: %s, Size: %d\n", *offset, p.Type, p.Size)
			fmt.Fprintf(os.Stdout, "  ObjRef: %s, ObjOffset: %d\n", p.HashOfRef, p.RefOffset)
			fmt.Fprintf(os.Stdout, "  Data(starts below):\n")
		}
		fmt.Fprintf(os.Stdout, "%s", p.Data)
		return
	}

	if *verifyPack {
		showVerifyPack(f, inIdx)
		return
	}

}
