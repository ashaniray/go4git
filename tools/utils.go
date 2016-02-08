package main

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

/////// Begin changes by Ashani ///////////////////

func GetArgInputFile() (*os.File, error) {
	args := flag.Args()
	if len(args) > 0 {
		return os.Open(args[0])
	} else {
		return os.Stdin, nil
	}
}

func GenSHA1(in *os.File, objType string) ([]byte, error) {
	buf := new(bytes.Buffer)
	length, err := buf.ReadFrom(in)
	if err != nil {
		return nil, err
	}
	prefix := fmt.Sprintf("%s %d", objType, length)

	h := sha1.New()
	io.WriteString(h, prefix)
	h.Write([]byte{0})
	buf.WriteTo(h)
	return h.Sum(nil), nil
}

func Unzlib(in *os.File, out *os.File) error {
	r, err := zlib.NewReader(in)
	if err != nil {
		return err
	}
	defer r.Close()
	_, err = io.Copy(out, r)
	return err
}

// readType is to be called by other read object functions
// So arg is buffer instead of file
func readTypeFromBuffer(buf *bytes.Buffer) (string, int, error) {
	//Read the type
	objType, err := buf.ReadString(' ')
	if err != nil {
		return "", 0, err
	}

	//Read the length of tree
	lenAsString, err := buf.ReadString(0)
	if err != nil {
		return "", 0, err
	}

	l, err := strconv.Atoi(lenAsString[:len(lenAsString)-1])
	return objType[:len(objType)-1], l, err
}

func ReadType(in *os.File) (string, int, error) {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(in)
	if err != nil {
		return "", 0, err
	}
	return readTypeFromBuffer(buf)
}

func ReadTree(in *os.File) (Tree, error) {

	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(in)
	if err != nil {
		return Tree{}, err
	}

	_, _, err = readTypeFromBuffer(buf)
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

///////////////////End changes by Ashani////////////
