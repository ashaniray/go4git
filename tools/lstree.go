package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := getArgInputFile()
	if err != nil {
		panic(err)
	}
	tree, err := readTree(f)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stdout, "%v\n", tree)
}

