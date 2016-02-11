package main

import (
	"flag"
	"fmt"
	"os"
)

var indexAt = flag.Int("i", 0, "index")
var hash = flag.String("h", "", "Hash of object")

func main() {
	flag.Parse()
	f, err := GetArgInputFile()
	var index PackIndex
	if len(*hash) > 0 {
		index, err = GetObjectForHash(*hash, f)
	} else {
		index, err = ReadPackIndexAt(*indexAt, f)
	}
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stdout, "%s\n", index)
}
