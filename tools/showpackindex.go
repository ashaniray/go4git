package main

import (
	"flag"
	"fmt"
	"github.com/ashaniray/go4git"
	"io"
	"os"
)

var size = flag.Bool("c", false, "Prints the number of index objects")
var hash = flag.String("h", "", "Hash of object to lookup")

func showSize(in io.ReadSeeker) {
	count, err := go4git.GetTotalCount(in)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stdout, "%d\n", count)
}

func showHashObject(in io.ReadSeeker) {
	index, err := go4git.GetObjectForHash(*hash, in)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stdout, "%s\n", index)
}

func showAllIndex(in io.ReadSeeker) {
	indices, err := go4git.GetAllPackedIndex(in)
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(indices); i++ {
		fmt.Fprintf(os.Stdout, "%s\n", indices[i])
	}
}

func main() {
	flag.Parse()
	f, err := go4git.GetArgInputFile()
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
