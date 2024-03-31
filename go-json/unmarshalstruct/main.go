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

	data := []byte(`{"name":"foo","age":42,"password":"top-secret!"}`)

	var e Employee
	if err := json.Unmarshal(data, &e); err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", e)
	// STOP OMIT
}
