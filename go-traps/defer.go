package main

import (
	"fmt"
	"log"
)

// START OMIT
func main() {
	fmt.Println("My favorite fruits:", favoriteFruits())
}

func favoriteFruits() (fruits []string) {
	defer log.Println("Found", len(fruits), "favorite fruits")

	// Here is some very complicated logic to determine my favorite fruits
	return []string{"🍎", "🍌", "🥝"}
}

// STOP OMIT
