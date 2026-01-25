package main

import "fmt"

func main() {
	// --- ARRAY EXAMPLE ---
	// Arrays always have equal len and cap.
	myArray := [10]int{1, 2, 3}
	fmt.Println("--- Array ---")
	fmt.Printf("Length: %d, Capacity: %d\n", len(myArray), cap(myArray))
	// Output: Length: 10, Capacity: 10 (even though only 3 are set, 7 are 0s)

	// --- SLICE EXAMPLE ---
	// We create a slice with a length of 3 but a capacity of 5.
	mySlice := make([]int, 3, 5)
	mySlice[0] = 10
	mySlice[1] = 20
	mySlice[2] = 30

	fmt.Println("\n--- Slice ---")
	fmt.Printf("Length: %d, Capacity: %d\n", len(mySlice), cap(mySlice))
	// Output: Length: 3, Capacity: 5

	// Adding an element increases length
	mySlice = append(mySlice, 40)
	fmt.Println("\n--- After Append ---")
	fmt.Printf("Length: %d, Capacity: %d\n", len(mySlice), cap(mySlice))
}
