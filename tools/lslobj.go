package main

import (
	"fmt"
	"os"
	"flag"
)

var repoRoot = flag.String("d", ".", "path to repository root")

func main() {
	flag.Parse()

	gitDir, err := GitDir(*repoRoot) 

	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err) 
		return
	}


	fmt.Fprintln(os.Stdout, "REPO:", gitDir)
}
