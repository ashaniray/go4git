package main

import (
	"flag"
	"fmt"
	"github.com/ashaniray/go4git"
	"os"
)

var repoRoot = flag.String("d", ".", "path to repository root")

func main() {
	flag.Parse()

	repo, err := go4git.NewRepository(*repoRoot)

	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
		return
	}

	refs, err := repo.References()

	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
		return
	}

	for _, ref := range refs {
		fmt.Fprintln(os.Stdout, ref)
	}
}
