package main

import (
	"os"
)

func main() {
	f, err := GetArgInputFile()
	if err != nil {
		panic(err)
	}
	err = Unzlib(f, os.Stdout)
	if err != nil {
		panic(err)
	}
}

