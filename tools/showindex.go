package main

import (
	"flag"
	"fmt"
	"os"
	"io"
)

var size = flag.Bool("c", false, "Prints the number of index objects")
var hash = flag.String("h", "", "Hash of object to lookup")

func showSize(in io.ReadSeeker) {
	count, err := GetTotalCount(in)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stdout, "%d\n", count)
}

func showHashObject(in io.ReadSeeker) {
	index, err := GetObjectForHash(*hash, in)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stdout, "%s\n", index)
}

func showAllIndex(in io.ReadSeeker) {
	count, err := GetTotalCount(in)
	if err != nil {
		panic(err)
	}
	for i := 0; i < int(count); i++ {
		index, err := ReadPackIndexAt(i, in)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(os.Stdout, "%s\n", index)
	}
}

func main() {
	flag.Parse()
	f, err := GetArgInputFile()
	if err != nil {
		panic(err)
	}
	if *size {
		showSize(f)
		return
	}
	if len(*hash) > 0 {
		showHashObject(f)
		return
	}
	showAllIndex(f)
}

