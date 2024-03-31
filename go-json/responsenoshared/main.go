package main

import (
	"encoding/json"
	"fmt"
)

// START HELPER OMIT
type Success[T any] struct {
	Data T `json:"data,omitempty"`
}

type Failure struct {
	Error string `json:"error,omitempty"`
}

type Response[T any] struct {
	Success[T]
	Failure
}

// STOP HELPER OMIT

func main() {
	// START OMIT
	data := []byte(`[
		{"data": "hello"},
		{"error": "Resource not found"}
	]`)

	var res []Response[string]
	if err := json.Unmarshal(data, &res); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", res)
	fmt.Printf("Data: %v\n", res[0].Data)
	fmt.Printf("Error: %v\n", res[1].Error)
	// STOP OMIT
}
