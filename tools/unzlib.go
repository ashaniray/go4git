package main

import (
	"flag"
	"os"
	"bufio"
)

func main() {
	flag.Parse()
	f, err := GetArgInputFile()
	if err != nil {
		panic(err)
	}
	err = Unzlib(bufio.NewReader(f), bufio.NewWriter(os.Stdout))
	if err != nil {
		panic(err)
	}
}
