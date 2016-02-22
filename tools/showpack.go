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

func showStats(stats map[int]int) {
	maxChain := 0
	for k, _ := range stats {
		if k > maxChain {
			maxChain = k
		}
	}
	fmt.Printf("non delta: %d objects\n", stats[0])
	for i := 1; i <= maxChain; i++ {
		fmt.Printf("chain length = %d: %d object", i, stats[i])
		if stats[i] > 1 {
			fmt.Printf("s")
		}
		fmt.Printf("\n")
	}
}

func showVerifyPack(inPack io.ReadSeeker, inIdx io.ReadSeeker) {
	indices, err := go4git.GetAllPackedIndex(inIdx)
	if err != nil {
		panic(err)
	}
	sort.Sort(go4git.ByOffset(indices))
	cnt := len(indices)

	chainLength := make(map[int]int)

	o, err := go4git.ReadPackedObjectAtOffset(int64(indices[0].Offset), inPack, inIdx)
	if err != nil {
		panic(err)
	}
	chainLength[o.RefLevel]++
	for i := 0; i < cnt-1; i++ {
		next, _ := go4git.ReadPackedObjectAtOffset(int64(indices[i+1].Offset), inPack, inIdx)
		if err != nil {
			panic(err)
		}
		chainLength[next.RefLevel]++
		str := fmt.Sprintf("%s %s %d %d %d", o.Hash, o.ActualType, o.Size, next.StartOffset-o.StartOffset, o.StartOffset)
		if o.Type == go4git.REF_DELTA || o.Type == go4git.OFS_DELTA {
			str += fmt.Sprintf(" %d %s", o.RefLevel, o.BaseHash)
		}
		fmt.Printf("%s\n", str)
		if o.Hash != indices[i].Hash {
			panic(fmt.Sprintf("Hash do not match for index %d offset %d: Expected: %s, Computed: %s",
				i,
				indices[i].Offset,
				indices[i].Hash,
				o.Hash,
			))
		}
		o = next
		if i == cnt-2 {
			end, _ := inPack.Seek(0, os.SEEK_END)
			end -= 20
			str := fmt.Sprintf("%s %s\t%d %d %d", next.Hash, next.ActualType, next.Size, end-next.StartOffset, next.StartOffset)
			if o.Type == go4git.REF_DELTA || o.Type == go4git.OFS_DELTA {
				str += fmt.Sprintf(" %d %s", o.RefLevel, o.BaseHash)
			}
			fmt.Printf("%s\n", str)
		}
	}
	showStats(chainLength)
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

	if *offset != -1 {
		p, err := go4git.ReadPackedObjectAtOffset(*offset, f, inIdx)
		if err != nil {
			panic(err)
		}
		if *verbose {
			fmt.Printf("Details of Object at offset [%d] \n", *offset)
			fmt.Printf(" Type: %s\n", p.Type)
			fmt.Printf(" HashOfRef: %s\n", p.HashOfRef)
			fmt.Printf(" RefOffset: %d\n", p.RefOffset)
			fmt.Printf(" Size: %d\n", p.Size)
			fmt.Printf(" StartOffset: %d\n", p.StartOffset)
			fmt.Printf(" ActualType: %s\n", p.ActualType)
			fmt.Printf(" Hash: %s\n", p.Hash)
			fmt.Printf(" RefLevel: %d\n", p.RefLevel)
			fmt.Printf(" BaseHash: %s\n", p.BaseHash)
			fmt.Printf(" ---Data(starts below):---\n")
		}
		fmt.Printf("%s", p.Data)
		return
	}

	if *verifyPack {
		showVerifyPack(f, inIdx)
		fmt.Printf("%s: ok\n", packFile)
		return
	}
}
