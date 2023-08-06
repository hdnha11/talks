package main

import "fmt"

func main() {
	// START OMIT
	type Point struct{ x, y int }

	points := []Point{
		{1, 2},
		{3, 4},
		{5, 6},
	}

	fmt.Println(points)

	// Convert to pointer slice

	var ptrs []*Point
	for _, pt := range points {
		ptrs = append(ptrs, &pt)
	}

	for _, pt := range ptrs {
		fmt.Println(*pt)
	}
	// STOP OMIT
}
