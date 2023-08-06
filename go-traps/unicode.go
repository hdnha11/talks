package main

import "fmt"

func main() {
	// START OMIT
	const placeOfInterest = `⌘`

	fmt.Printf("plain string: ")
	fmt.Printf("%s", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("quoted string: ")
	fmt.Printf("%+q", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("hex bytes: ")
	fmt.Printf("% x", placeOfInterest)
	fmt.Printf("\n")
	// STOP OMIT
}
