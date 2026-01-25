package main

import "fmt"

func main() {
	names := make([]string, 3, 5)

	fmt.Println("Initial names slice:", names)
	fmt.Printf("Length: %d, Capacity: %d\n", len(names), cap(names))

	// Assign values to the slice
	names[0] = "Alice"
	names[1] = "Bob"
	names[2] = "Charlie"

	fmt.Println("Names slice:", names)
	fmt.Printf("Length: %d, Capacity: %d\n", len(names), cap(names))
}
