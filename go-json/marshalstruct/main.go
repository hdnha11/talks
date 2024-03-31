package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// START OMIT
	type Employee struct {
		Name    string `json:"name"`
		Age     int    `json:"age"`
		Address string `json:"address,omitempty"`

		password string // Unexported fields are ignored
	}

	e := Employee{
		Name:     "foo",
		Age:      42,
		password: "top-secret!",
	}

	data, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", data)
	// STOP OMIT
}
