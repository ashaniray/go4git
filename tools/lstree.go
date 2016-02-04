package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := GetArgInputFile()
	if err != nil {
		panic(err)
	}
	tree, err := ReadTree(f)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stdout, "%s\n", tree)
}

