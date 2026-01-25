package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	//Time Complexity: O(n), where n is the number of elements in the array. We visit each element exactly once.

	//Space Complexity: O(1), because we only use two extra variables regardless of how large the input array is.
	maxCount := 0
	currentCount := 0

	for _, num := range nums {
		if num == 1 {
			currentCount++
			if currentCount > maxCount {
				maxCount = currentCount
			}
		} else {
			currentCount = 0
		}
	}

	return maxCount
}

func main() {
	nums := []int{1, 1, 0, 1, 1, 1}
	result := findMaxConsecutiveOnes(nums)
	fmt.Println("Maximum consecutive ones:", result) // Output: 3
}
