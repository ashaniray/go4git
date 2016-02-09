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
	t, s, buff, err := ReadPackedObjectAtOffset(*offset, f)
	if err != nil {
		panic(err)
	}
	if *verbose {
		fmt.Fprintf(os.Stdout, "Object at offset [%d] => Type: %s, Size: %d, Data(starts below):\n", *offset, t, s)
	}
	fmt.Fprintf(os.Stdout, "%s", buff)
}

