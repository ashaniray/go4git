package main

import(
	"fmt"
	"os"
	"flag"
)

var repoRoot = flag.String("d", ".", "path to repo")


type Commit struct {
	Tree string
	Parent string
	Author string
	Committer string
}



func ParseCommit(sha string, root string) (*Commit, error) {

	commit := new(Commit)

	cmtObjPath, err := GetObjPath(sha, root)

	if err != nil {
		return nil, err
	}

	_, err = os.Open(cmtObjPath)

	if err != nil {
		return nil, err
	}

	return commit, nil
}

func main() {
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "ERROR:", "Invalid agruments")
		return
	}


	commit, err := ParseCommit(flag.Arg(0), *repoRoot)

	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err) 
		return
	}

	fmt.Println(commit)
}
