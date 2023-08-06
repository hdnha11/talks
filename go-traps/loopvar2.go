package main

import (
	"fmt"
	"time"
)

func main() {
	// START OMIT
	fruits := []string{"Apple", "Banana", "Kiwi"}

	for _, fruit := range fruits {
		go func() {
			fmt.Printf("YUM YUM! %s\n", fruit)
		}()
	}
	// STOP OMIT

	time.Sleep(1 * time.Second)
}
