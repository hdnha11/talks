package main

import "fmt"

func main() {
	// START OMIT
	m := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
	}

	// Fetching
	fmt.Println(m["four"])

	// Traversal
	for name, value := range m {
		fmt.Println(name, "is english for number", value)
	}
	// STOP OMIT
}
