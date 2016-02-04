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
	t, _, err := ReadType(f)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stdout, "Type:%s\n", t)
}

