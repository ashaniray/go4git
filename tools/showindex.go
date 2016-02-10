package main

import (
	"flag"
	"fmt"
	"os"
)

var indexAt = flag.Int("i", 0, "index")

func main() {
	flag.Parse()
	f, err := GetArgInputFile()

	index, err := ReadPackIndexAt(*indexAt, f)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stdout, "%s\n", index)
}

