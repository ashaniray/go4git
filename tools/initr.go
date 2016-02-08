package main

import (
	"flag"
	"fmt"
	"os"
	"path"
)

var isBare = flag.Bool("bare", false, "create a bare repository instead")

func main() {
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "ERROR: Invalid argument.")
		return
	}

	repoRoot := flag.Arg(0)
	err := CreateRepository(repoRoot, *isBare)

	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
	} else {
		fmt.Fprintf(os.Stdout, "created empty repository %s\n", path.Base(repoRoot))
	}
}
