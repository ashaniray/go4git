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
	_, t, err := go4git.ReadObjectType(f)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stdout, "Type:%s\n", t)
}
