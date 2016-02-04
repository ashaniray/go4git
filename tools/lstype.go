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
	t, _, err := readType(f)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stdout, "Type:%s\n", t)
}

