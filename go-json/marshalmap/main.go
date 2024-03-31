package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// START OMIT
	m := map[string]any{
		"name": "foo",
		"age":  42,
	}

	data, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", data)
	// STOP OMIT
}
