package main

import "fmt"

func main() {
	// START OMIT
	str := "🙌 Xin chào, Nhã!"

	for i := 0; i < len(str); i++ {
		fmt.Printf("Character #%02d is %c\n", i, str[i])
	}
	// STOP OMIT
}
