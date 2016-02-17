package main

import (
	"flag"
	"fmt"
	"github.com/ashaniray/go4git"
	"os"
)

func main() {
	flag.Parse()
	f, err := go4git.GetArgInputFile()
	if err != nil {
		panic(err)
	}
	tree, err := go4git.ParseTree(f)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stdout, "%s\n", tree)
}
