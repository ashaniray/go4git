package main

import (
	"flag"
	"github.com/ashaniray/go4git"
	"os"
)

func main() {
	flag.Parse()
	f, err := go4git.GetArgInputFile()
	if err != nil {
		panic(err)
	}
	err = go4git.Zlib(f, os.Stdout)
	if err != nil {
		panic(err)
	}
}
