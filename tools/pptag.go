package main

import (
	"fmt"
	"os"
	"flag"
	"github.com/ashaniray/go4git"
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
