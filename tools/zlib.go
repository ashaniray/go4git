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
	err = Zlib(f, os.Stdout)
	if err != nil {
		panic(err)
	}
}
