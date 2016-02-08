package main

import (
	"flag"
	"fmt"
	"os"
)

var repoRoot = flag.String("d", ".", "path to repository root")

func main() {
	flag.Parse()

	objs, err := AllObjects(*repoRoot)

	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
	}

	for _, obj := range objs {
		fmt.Fprintln(os.Stdout, obj)
	}
}
