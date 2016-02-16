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

	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
		return
	}

	tag, err := go4git.ParseTag(f)

	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
	}

	fmt.Fprintln(os.Stdout, tag)
}
