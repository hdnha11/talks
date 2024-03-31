package main

import (
	"encoding/json"
	"fmt"
)

// START HELPER OMIT
type MagicSlice[T any] []T

func (s *MagicSlice[T]) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}

	if data[0] == '[' {
		var v []T
		err := json.Unmarshal(data, &v)
		*s = MagicSlice[T](v)
		return err
	}

	var v T
	err := json.Unmarshal(data, &v)
	*s = MagicSlice[T]{v}
	return err
}

// STOP HELPER OMIT

func main() {
	// START OMIT
	data := []byte(`["one", ["two", "three"]]`)

	var s []MagicSlice[string]
	if err := json.Unmarshal(data, &s); err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", s)
	// STOP OMIT
}
