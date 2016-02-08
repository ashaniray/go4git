package main

import(
	"fmt"
	"os"
	"flag"
	"io"
)

var repoRoot = flag.String("d", ".", "path to repo")


type Commit struct {
	Tree string
	Parent string
	Author string
	Committer string
}



func ParseCommit(in io.Reader) (*Commit, error) {

	commit := new(Commit)

	return commit, nil
}

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
