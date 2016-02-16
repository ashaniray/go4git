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

	cmt, err := go4git.ParseCommit(f)

	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
		return
	}

	fmt.Fprintln(os.Stdout, cmt)
}
