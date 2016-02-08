package main

import (
	"flag"
	"fmt"
	"os"
)

var objType = flag.String("t", "blob", "The type of object, e.g. tree, blob, commit, or tag")

func main() {
	flag.Parse()
	f, err := GetArgInputFile()
	if err != nil {
		panic(err)
	}
	b, err := GenSHA1(f, *objType)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stdout, "%s", b)
}
