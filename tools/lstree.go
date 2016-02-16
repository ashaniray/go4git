package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/ashaniray/go4git"
)

func main() {
	flag.Parse()
	f, err := go4git.GetArgInputFile()
	if err != nil {
		panic(err)
	}
	tree, err := go4git.ReadTree(f)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stdout, "%s\n", tree)
}
