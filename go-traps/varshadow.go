package main

import (
	"log"
)

func main() {
	// START OMIT
	var message string

	if message, err := hello(); err != nil {
		log.Println(err, message)
	}

	log.Printf("The message is: %s", message)
	// STOP OMIT
}

func hello() (string, error) {
	return "Xin chào!", nil
}
