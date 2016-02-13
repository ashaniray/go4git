package main

import (
	"flag"
	"fmt"
	"os"
)

var repoRoot = flag.String("d", ".", "path to a git repo")

func main() {
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "ERROR:", "Illegal arguments")
		return
	}

	repo, err := NewRepository(*repoRoot)

	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
		return
	}

	startSha := flag.Arg(0)

	c, err := repo.LookupCommit(startSha)

	if err != nil  {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
		return
	}

	for {
		fmt.Fprintln(os.Stdout, c.Id, c.Message)

		if !c.HasParent() {
			break
		}

		c, err = repo.LookupCommit(c.Parent)

		if err != nil {
			fmt.Fprintln(os.Stderr, "ERROR:", err)
			break
		}
	}

}
