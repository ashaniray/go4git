package main

import (
	"flag"
	"fmt"
	"os"
)

var offset = flag.Int64("s", -1, "The offset to read from the pack file")
var infoOnly = flag.Bool("t", false, "Output verbose information")

func main() {
	flag.Parse()
	f, err := GetArgInputFile()
	if err != nil {
		panic(err)
	}
	t, s, buff, err := ReadPackedDataAtOffset(*offset, f)
	if err != nil {
		panic(err)
	}
	if *infoOnly {
		fmt.Fprintf(os.Stdout, "Type: %s, Size: %d\n", t, s)
	} else {
		fmt.Fprintf(os.Stdout, "%s", buff)
	}
}

