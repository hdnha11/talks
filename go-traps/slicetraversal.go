package main

import "fmt"

func main() {
	// START OMIT
	languages := []string{"C", "Java"}
	languages = append(languages, "Go")

	for lang := range languages {
		fmt.Println("Hello", lang)
	}
	// STOP OMIT
}
