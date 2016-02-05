package main

import (
	"os"
	"fmt"
	"flag"
)


var repoRoot = flag.String("d", ".", "path to repository root")

func main() {
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "ERROR: Invalid arguments")
		return
	}

	p, err := GetObjPath(flag.Arg(0), *repoRoot)

	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
	} else {
		fmt.Fprintf(os.Stdout, p)
	}

}
