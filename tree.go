package go4git

import (
	"bytes"
	"fmt"
	"os"
)

type TreeItem struct {
	isBlob bool
	mode   string
	name   string
	hash   []byte
}

type Tree struct {
	items []TreeItem
}

func (item TreeItem) String() string {
	t := "blob"
	if item.isBlob == false {
		t = "tree"
	}
	var data bytes.Buffer
	for _, dataByte := range item.hash {
		data.WriteString(fmt.Sprintf("%.2x", dataByte))
	}
	return fmt.Sprintf("{%s %s %s %s}", t, item.mode, item.name, data.String())
}

func (tree Tree) String() string {
	var treeAsString bytes.Buffer
	treeAsString.WriteString("[\n")
	for _, item := range tree.items {
		treeAsString.WriteString(fmt.Sprintf("%s\n", item))
	}
	treeAsString.WriteString("]")
	return treeAsString.String()
}

func ParseTree(in *os.File) (Tree, error) {

	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(in)
	if err != nil {
		return Tree{}, err
	}

	_, _, err = parseHeader(buf)
	if err != nil {
		return Tree{}, err
	}

	items := []TreeItem{}
	for buf.Len() > 0 {
		//Read the mode
		mode, err := buf.ReadString(' ')
		if err != nil {
			return Tree{}, err
		}
		mode = mode[:len(mode)-1]

		isBlob := false
		if mode[0] == '1' {
			isBlob = true
		}

		//Read the name
		name, err := buf.ReadString(0)
		if err != nil {
			return Tree{}, err
		}

		//Read the 20 byte hash
		hash := buf.Next(20)

		treeItem := TreeItem{isBlob: isBlob, mode: mode, hash: hash, name: name}
		items = append(items, treeItem)
	}

	return Tree{items: items}, nil
}
