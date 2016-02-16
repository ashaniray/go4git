package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/ashaniray/go4git"
)

var repoRoot = flag.String("d", ".", "path to repository root")

func main() {
	flag.Parse()

	repo, err := go4git.NewRepository(*repoRoot)

	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
		return
	}

	objs, err := repo.LooseObjects()

	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
		return
	}

	for _, obj := range objs {
		fmt.Fprintln(os.Stdout, obj)
	}
}
