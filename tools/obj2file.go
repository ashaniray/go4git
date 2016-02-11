package main

import (
	"flag"
	"fmt"
	"os"
)

var repoRoot = flag.String("d", ".", "path to repository root")

func main() {
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "ERROR: Invalid arguments")
		return
	}

	repo, err := NewRepository(*repoRoot)

	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
		return
	}
    
	fmt.Fprintf(os.Stdout, repo.LooseObjPath(flag.Arg(0)))
}
