package main

import (
	"os"
	"fmt"
)

func main() {
	f, err := getArgInputFile()
	if err != nil {
		panic(err)
	}
	b, err := genSHA1(f)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stdout, "%s", b)
}

