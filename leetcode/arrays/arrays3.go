package main

import "fmt"

func main() {
	// 1. SETTING/GETTING VALUES
	// Arrays are 0-indexed.
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"

	fmt.Println("First color:", colors[0])

	// 2. FINDING LENGTH
	// Use the built-in len() function.
	fmt.Println("Array length:", len(colors))

	// 3. ITERATING (The 'range' keyword)
	// This is the most idiomatic way to loop in Go.
	fmt.Println("Looping through elements:")
	for index, value := range colors {
		fmt.Printf("Index %d has value: %s\n", index, value)
	}

	// 4. MULTI-DIMENSIONAL ARRAYS
	// Useful for grids or matrices.
	var matrix [2][2]int
	matrix[0] = [2]int{1, 2}
	matrix[1] = [2]int{3, 4}
	fmt.Println("2D Matrix:", matrix)
}
