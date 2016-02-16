package main

import (
	"bufio"
	"flag"
	"os"
	"github.com/ashaniray/go4git"
)

func main() {
	flag.Parse()
	f, err := go4git.GetArgInputFile()
	if err != nil {
		panic(err)
	}
	err = go4git.Unzlib(bufio.NewReader(f), bufio.NewWriter(os.Stdout))
	if err != nil {
		panic(err)
	}
}
