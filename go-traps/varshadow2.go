package main

import "fmt"

// START OMIT
func main() {
	var x Animal = Dog{"Rosie"}

	if x, ok := x.(Human); ok {
		fmt.Println(x.lastName, "doesn't want to be treated like dogs.")
	} else {
		fmt.Println(x.Say())
	}
}

type Animal interface{ Say() string }

type Dog struct{ name string }

func (d Dog) Say() string { return fmt.Sprintf("%s barks", d.name) }

type Human struct{ firstName, lastName string }

// Humans are technically animals, and they say things. OMIT
func (h Human) Say() string { return fmt.Sprintf("%s %s speaks", h.firstName, h.lastName) }

// STOP OMIT
