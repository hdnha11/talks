package main

import "fmt"

// START OMIT
type Animal interface {
	Say()
	Greet()
}

type Cat struct{}

func (c Cat) Say() { fmt.Println("meow") }
func (c Cat) Greet() {
	fmt.Println("I'm a cat!")
	c.Say()
}

type Tiger struct{ Cat }

func (t Tiger) Say() { fmt.Println("roar") }

func main() {
	var x Animal = Tiger{}
	x.Greet()
}

// STOP OMIT
