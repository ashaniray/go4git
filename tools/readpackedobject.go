package main

import (
	"flag"
	"fmt"
	"os"
	"io"
)

var offset = flag.Int64("s", -1, "The offset to read from the pack file")
var verbose = flag.Bool("t", false, "Output verbose information")
var verify = flag.Bool("v", true, "Produce output of git pack-verify -v")

func showVerifyPack(in io.ReadSeeker) {
	var i int64 = 12
	for {
		p, err := ReadPackedObjectAtOffset(i, in)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(os.Stdout, "%s\n", p)
		break
	}
}

func main() {
	flag.Parse()
	f, err := GetArgInputFile()
	if err != nil {
		panic(err)
	}
	showVerifyPack(f)
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
