package main

import "fmt"

func main() {
	arr := []int{10, 20, 30}
	fmt.Println("Original array:", arr)

	// 1. End
	arr = append(arr, 40)
	fmt.Println("After appending 40 at the end:", arr)

	// 2. Beginning
	arr = append([]int{0}, arr...)
	fmt.Println("After inserting 0 at the beginning:", arr)

	// 3. Middle (insert 25 at index 3)
	index := 3
	arr = append(arr[:index], append([]int{25}, arr[index:]...)...)
	fmt.Println("After inserting 25 at index 3:", arr)

	// 4. Remove element at index 2
	indexToRemove := 2
	arr = append(arr[:indexToRemove], arr[indexToRemove+1:]...)
	fmt.Println("After removing element at index 2:", arr)
}
