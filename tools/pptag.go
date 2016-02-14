package main

import (
	"fmt"
	"os"
	"flag"
)

func main() {
	flag.Parse()

	f, err := GetArgInputFile()

	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)	
		return
	}

	tag, err := parseTag(f) 

	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)	
	}

	fmt.Fprintln(os.Stdout, tag)
}
