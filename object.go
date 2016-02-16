package go4git

// holds all functions common to all object types

import (
	"bytes"
	"strconv"
	"strings"
)

func parseHeader(buff *bytes.Buffer) (int, string, error) {
	header, err := buff.ReadString(0)

	if err != nil {
		return 0, "", err
	}

	xs := strings.Split(header, " ")
	objType, objSize := xs[0], xs[1]

	objSize = objSize[:len(objSize)-1] // remove trailing null

	size, err := strconv.Atoi(objSize)

	if err != nil {
		return 0, objType, err
	}

	return size, objType, nil
}
