package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// START HELPER OMIT
type MagicInt int

func (i *MagicInt) UnmarshalJSON(data []byte) error {
	var v int
	if err := json.Unmarshal(bytes.Trim(data, `"`), &v); err != nil {
		return err
	}

	*i = MagicInt(v)
	return nil
}

// STOP HELPER OMIT

func main() {
	// START OMIT
	data := []byte(`["42",123]`)
	var scores []MagicInt
	if err := json.Unmarshal(data, &scores); err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", scores)
	// STOP OMIT
}
