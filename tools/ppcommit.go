package main

import(
	"fmt"
	"os"
	"flag"
)

func main() {
	flag.Parse()

	f, err := GetArgInputFile()

	cmt, err := ParseCommit(f)

	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)	
		return
	}

	fmt.Fprintln(os.Stdout, cmt)
}
