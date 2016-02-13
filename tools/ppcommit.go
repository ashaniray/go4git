package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()

	f, err := GetArgInputFile()

	cmt, err := parseCommit(f)

	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
		return
	}

	fmt.Fprintln(os.Stdout, cmt)
}
