package main

import (
	"fmt"
	"bytes"
)

type TreeItem struct {
	isBlob bool
	mode string
	name string
	hash []byte
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

