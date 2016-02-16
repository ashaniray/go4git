package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/ashaniray/go4git"
)

func main() {
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "ERROR:", "Invalid arguments")
		return
	}

	repo, err := go4git.NewRepository(flag.Arg(0))

	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
		return
	}

	fmt.Fprintf(os.Stdout, "Bare:     %t\n", repo.IsBare())
	fmt.Fprintf(os.Stdout, "Path:     %s\n", repo.Path())
	fmt.Fprintf(os.Stdout, "Work Dir: %s\n", repo.WorkDir())

}
