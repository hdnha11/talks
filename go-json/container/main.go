package main

import (
	"encoding/json"
	"fmt"
)

// START TYPE OMIT
type Circle struct {
	Radius float64 `json:"radius"`
}

type Rectangle struct {
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

type Triangle struct {
	Base   float64 `json:"base"`
	Height float64 `json:"height"`
}

// STOP TYPE OMIT
// START CONTAINER PART1 OMIT
type Result struct {
	Shape any
}

func (r *Result) UnmarshalJSON(data []byte) error {
	var header struct {
		Shape string `json:"shape"`
	}
	if err := json.Unmarshal(data, &header); err != nil {
		return err
	}

	// STOP CONTAINER PART1 OMIT
	// START CONTAINER PART2 OMIT
	var err error
	switch header.Shape {
	case "circle":
		target := Circle{}
		err = json.Unmarshal(data, &target)
		r.Shape = target
	case "rectangle":
		target := Rectangle{}
		err = json.Unmarshal(data, &target)
		r.Shape = target
	case "triangle":
		target := Triangle{}
		err = json.Unmarshal(data, &target)
		r.Shape = target
	}
	if err != nil {
		return err
	}

	return nil
}

// STOP CONTAINER PART2 OMIT

func main() {
	// START OMIT
	data := []byte(`[
		{"shape": "circle", "radius": 3.14},
		{"shape": "rectangle", "width": 7.2, "height": 4},
		{"shape": "triangle", "base": 8.6, "height": 6.2}
	]`)

	var result []Result
	if err := json.Unmarshal(data, &result); err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", result)
	// STOP OMIT
}
