package main

import "fmt"

func main() {
	// Method A: Declare then assign
	var a [5]int
	fmt.Println("Empty array:", a) // Output: [0 0 0 0 0]

	// Method B: Array Literal
	b := [5]int{10, 20, 30, 40, 50}
	fmt.Println("Initialized array:", b)

	// Method C: Let Go count the elements for you
	c := [...]int{1, 2, 3}
	fmt.Println("Compiler-counted size:", len(c))
}
