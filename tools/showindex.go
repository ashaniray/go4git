package main

import (
	"flag"
	"fmt"
	"github.com/ashaniray/go4git"
	"io"
	"os"
)

func showIndex(in io.ReadSeeker) {
	index, err := go4git.ParseIndex(in)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stdout, "%s", index)
}

func main() {
	flag.Parse()
	f, err := go4git.GetArgInputFile()
	if err != nil {
		panic(err)
	}
	showIndex(f)
	return
}
