package main

import (
	"bufio"
	"flag"
	"os"
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
