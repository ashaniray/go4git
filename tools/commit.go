package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
)

type CommitFields map[string]string

type Commit struct {
	Id        string
	Tree      string
	Parent    string
	Author    *Person
	Committer *Person
	Message   string
}

func (c *Commit) HasParent() bool {
	return len(c.Parent) != 0
}

func (c *Commit) String() string {
	buff := new(bytes.Buffer)
	fmt.Fprintf(buff, "Tree:      %s\n", c.Tree)
	fmt.Fprintf(buff, "Parent:    %s\n", c.Parent)
	fmt.Fprintf(buff, "Author:    %s\n", c.Author)
	fmt.Fprintf(buff, "Committer: %s\n", c.Committer)
	fmt.Fprintf(buff, "Message:   %s\n", c.Message)

	return string(buff.Bytes())
}

func kvPair(s string) (string, string) {
	xs := strings.Split(s, " ")
	return xs[0], strings.Join(xs[1:], " ")
}

func parseCommitBody(buff *bytes.Buffer, size int) (CommitFields, error) {
	body := string(buff.Next(size))

	lines := strings.Split(body, "\n")

	fields := make(map[string]string)

	// TODO: add validations and return error
	for i, l := range lines {
		if len(l) == 0 {
			fields["message"] = strings.Trim(strings.Join(lines[i:], "\n"), "\n")
			break
		} else {
			k, v := kvPair(l)
			fields[k] = v
		}
	}

	return fields, nil
}

func (cf CommitFields) ToCommit() *Commit {
	commit := new(Commit)

	commit.Tree = cf["tree"]
	commit.Author = parsePerson(cf["author"])
	commit.Committer = parsePerson(cf["committer"])
	commit.Message = cf["message"]

	parent, ok := cf["parent"]

	if ok {
		commit.Parent = parent
	}
	return commit
}

func parseCommit(in io.Reader) (*Commit, error) {

	buff := new(bytes.Buffer)
	buff.ReadFrom(in)

	size, typ, err := parseHeader(buff)

	if err != nil {
		return nil, err
	}

	if typ != "commit" {
		return nil, errors.New("not a commit")
	}

	body, err := parseCommitBody(buff, size)

	if err != nil {
		return nil, err
	}

	return body.ToCommit(), nil
}
