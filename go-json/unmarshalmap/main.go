package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// START OMIT
	data := []byte(`{"name":"foo","age":42}`)

	var m any
	if err := json.Unmarshal(data, &m); err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", m)
	// STOP OMIT
}
