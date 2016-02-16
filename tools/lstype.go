package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/ashaniray/go4git"
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
