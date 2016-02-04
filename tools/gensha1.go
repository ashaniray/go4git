package main

import (
	"os"
	"fmt"
)

func main() {
	f, err := GetArgInputFile()
	if err != nil {
		panic(err)
	}
	b, err := GenSHA1(f)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stdout, "%s", b)
}

