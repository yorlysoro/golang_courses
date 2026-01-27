package main

// sortedSquares takes a sorted integer array and return a new array of squares
// Time Complexity: O(n) - We traverse the array exactly once.
// Space Complexity: O(n) - We create a new array to store the squares.
func sortedSquares(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	left, right := 0, n-1
	position := n - 1

	// Fill the result array from back to front
	for left <= right {
		leftSquare := nums[left] * nums[left]
		rightSquare := nums[right] * nums[right]

		if leftSquare > rightSquare {
			result[position] = leftSquare
			left++
		} else {
			result[position] = rightSquare
			right--
		}
		position--
	}

	return result
}

func main() {
	// Test Case 1: Standard mixed numbers
	nums1 := []int{-4, -1, 0, 3, 10}
	squared1 := sortedSquares(nums1)
	println("Test Case 1 Result:")
	for _, v := range squared1 {
		print(v, " ")
	}
	println()

	// Test Case 2: Negative and positive numbers
	nums2 := []int{-7, -3, 2, 3, 11}
	squared2 := sortedSquares(nums2)
	println("Test Case 2 Result:")
	for _, v := range squared2 {
		print(v, " ")
	}
	println()

	// Test Case 3: All negative numbers
	nums3 := []int{-5, -3, -1}
	squared3 := sortedSquares(nums3)
	println("Test Case 3 Result:")
	for _, v := range squared3 {
		print(v, " ")
	}
	println()

	// Test Case 4: Single element (Constraint: 1 <= nums.length )
	nums4 := []int{-2}
	squared4 := sortedSquares(nums4)
	println("Test Case 4 Result:")
	for _, v := range squared4 {
		print(v, " ")
	}
	println()
}
