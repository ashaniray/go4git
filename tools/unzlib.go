package main

import (
	"flag"
	"os"
)

func main() {
	flag.Parse()
	f, err := GetArgInputFile()
	if err != nil {
		panic(err)
	}
	err = Unzlib(f, os.Stdout)
	if err != nil {
		panic(err)
	}
}
