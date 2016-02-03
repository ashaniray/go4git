package main

import (
	"os"
)

func main() {
	f, err := getArgInputFile()
	if err != nil {
		panic(err)
	}
	err = genSHA1(f, os.Stdout)
	if err != nil {
		panic(err)
	}
}

