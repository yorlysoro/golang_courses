package main

import "fmt"

// findNumbers returns the count of numbers that have an even number of digits.
func findNumbers(nums []int) int {
	// Time Complexity: O(N×K), where N is the number of elements and K is the maximum number of digits (here, K is small, max 6). Effectively O(N).

	// Space Complexity: O(1), as we only use a few integer variables regardless of input size.
	totalEven := 0

	for _, n := range nums {
		if hasEvenDigits(n) {
			totalEven++
		}
	}

	return totalEven
}

// hasEvenDigits is a helper function to determine digit parity.
func hasEvenDigits(num int) bool {
	count := 0
	// Standard math approach: divide by 10 to strip digits
	for num > 0 {
		num /= 10
		count++
	}
	// Check if count is even
	return count%2 == 0
}

func main() {
	example1 := []int{12, 345, 2, 6, 7896}
	fmt.Printf("Example 1 Result: %d\n", findNumbers(example1))
}
