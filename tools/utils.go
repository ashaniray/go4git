package main

import (
	"compress/zlib"
	"io"
	"os"
)

func getArgInputFile() (*os.File, error) {
	args := os.Args[1:]
	if len(args) > 0 {
		return os.Open(args[0])
	} else {
		return os.Stdin, nil
	}
}

func unzlib(in *os.File, out *os.File) error {
	r, err := zlib.NewReader(in)
	if err != nil {
		return err
	}
	defer r.Close()
	_, err = io.Copy(out, r)
	return err
}

