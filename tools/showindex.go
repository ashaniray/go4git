package main

import (
	"flag"
	"fmt"
	"os"
)

var indexAt = flag.Int("i", 0, "index")
var hash = flag.String("h", "", "Hash of object")
var output = flag.String("o", "all", "Ouputs a specific field: "+
	"\"hash\", \"offset\", \"crc\"")

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
	switch *output {
	case "hash":
		fmt.Fprintf(os.Stdout, "%s\n", index.hash)
	case "offset":
		fmt.Fprintf(os.Stdout, "%d\n", index.offset)
	case "crc":
		fmt.Fprintf(os.Stdout, "%s\n", index.crc)
	default:
		fmt.Fprintf(os.Stdout, "%s\n", index)
	}
}

