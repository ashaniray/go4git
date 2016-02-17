package main

import (
	"flag"
	"fmt"
	"github.com/ashaniray/go4git"
	"os"
)

var (
	objType = flag.String("c", "all", "print fixture for object type. [all|commit|blob|tree|tag]")
	objSize = flag.String("s", "sm", "print fixture of given size. [xs|sm|md|lg|xl]") // TODO: later ...
)

var m = map[string][]byte{
	"commit": go4git.COMMIT_DATA,
	"blob":   go4git.BLOB_DATA,
	"tree":   go4git.TREE_DATA,
	"tag":    go4git.TAG_DATA,
}

func printAll() {
	for k, v := range m {
		fmt.Fprintln(os.Stdout, k, "[")
		os.Stdout.Write(v)
		fmt.Fprintln(os.Stdout, "\n]")
	}
}

func printFixture(t string, s string) {

	v, ok := m[t]
	if ok {
		os.Stdout.Write(v)
	} else {
		fmt.Fprintln(os.Stderr, "ERROR: Invalid object type or size.")
	}
}

func main() {
	flag.Parse()

	if *objType == "all" {
		printAll()
		return
	}

	printFixture(*objType, *objSize)
}
