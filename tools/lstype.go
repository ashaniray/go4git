package main

import (
	"flag"
	"fmt"
	"github.com/ashaniray/go4git"
	"os"
)

func main() {
	flag.Parse()
	f, err := go4git.GetArgInputFile()
	if err != nil {
		panic(err)
	}
	t, _, err := go4git.ReadType(f)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stdout, "Type:%s\n", t)
}
