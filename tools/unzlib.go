package main

import (
	"compress/zlib"
	"io"
	"os"
)

func main() {
	r, err := zlib.NewReader(os.Stdin)
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, r)
	r.Close()
}

