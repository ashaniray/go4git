package main

import (
	"flag"
	"fmt"
	"os"
)

var offset = flag.Int64("s", -1, "The offset to read from the pack file")
var verbose = flag.Bool("t", false, "Output verbose information")

func main() {
	flag.Parse()
	f, err := GetArgInputFile()
	if err != nil {
		panic(err)
	}
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
