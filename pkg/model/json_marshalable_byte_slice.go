package model

import "fmt"

type JSONMarshalableByteSlice []byte

func (j JSONMarshalableByteSlice) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", string(j))), nil
}
