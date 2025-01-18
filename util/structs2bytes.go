package util

import (
	"bytes"
	"encoding/gob"
)

func StructsToBytes[T any](structs []T) []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	err := enc.Encode(structs)
	if err != nil {
		panic(err)
	}

	return buf.Bytes()
}