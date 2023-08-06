package main

import (
	"fmt"
)

// START READ OMIT
type DivByZeroError struct {
	numerator int
}

func (e *DivByZeroError) Error() string {
	return fmt.Sprintf(
		"division by zero error: cannot divide the number %d by zero", e.numerator)
}

// STOP READ OMIT

// START EDIT OMIT
func main() {
	n, err := calc(42, 2)
	if err != nil {
		fmt.Println("Oops:", err)
		return
	}

	fmt.Println("Result:", n)
}

func div(a, b int) (int, *DivByZeroError) {
	if b == 0 {
		return 0, &DivByZeroError{a}
	}

	return a / b, nil
}

func calc(a, b int) (int, error) {
	return div(a, b)
}

// STOP EDIT OMIT
