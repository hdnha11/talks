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
	Status string `json:"status"`
	Success[T]
	Failure
}

// STOP HELPER OMIT

func main() {
	// START OMIT
	data := []byte(`[
		{"status": "ok", "data": "hello"},
		{"status": "error", "error": "Resource not found"}
	]`)

	var res []Response[string]
	if err := json.Unmarshal(data, &res); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", res)
	fmt.Printf("%+v\n", res[0].Status)
	fmt.Printf("%+v\n", res[1].Status)
	// STOP OMIT
}
