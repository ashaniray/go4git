package main

type TreeItem struct {
	isBlob bool
	mode string
	name string
	hash []byte
}


type Tree struct {
	items []TreeItem
}
