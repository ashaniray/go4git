package main

import (
	"os"
)

func main() {
	f, err := getArgInputFile()
	if err != nil {
		panic(err)
	}
	err = unzlib(f, os.Stdout)
	if err != nil {
		panic(err)
	}
}

