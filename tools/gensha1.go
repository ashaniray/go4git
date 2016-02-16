package main

import (
	"flag"
	"fmt"
	"github.com/ashaniray/go4git"
	"os"
)

var objType = flag.String("t", "blob", "The type of object, e.g. tree, blob, commit, or tag")

func main() {
	flag.Parse()
	f, err := go4git.GetArgInputFile()
	if err != nil {
		panic(err)
	}
	b, err := go4git.GenSHA1(f, *objType)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stdout, "%s", b)
}
