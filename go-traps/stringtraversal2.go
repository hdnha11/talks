package main

import "fmt"

func main() {
	// START OMIT
	str := "🙌 Xin chào, Nhã!"

	for i, c := range str {
		fmt.Printf("Character #%02d is %c\n", i, c)
	}
	// STOP OMIT
}
