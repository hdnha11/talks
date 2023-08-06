package main

import (
	"fmt"
)

// START OMIT
func main() {
	s := Student{id: "stu0001"}
	defer s.print()
	s.id = "stu0002"
}

type Student struct {
	id string
}

func (s Student) print() {
	fmt.Println(s.id)
}

// STOP OMIT
