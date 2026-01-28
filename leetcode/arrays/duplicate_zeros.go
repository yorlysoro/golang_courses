package main

// duplicateZeros modifies the input array in-place by duplicating each zero
// and shifting subsequent elements to the right.

func duplicateZeros(arr []int) {
	n := len(arr)
	zerosToDuplicate := 0
	lastIndex := n - 1

	// First pass: Count the number of zeros to be duplicated
	// and the boundary of the elements that will fit in the final array.
	for i := 0; i <= lastIndex-zerosToDuplicate; i++ {
		if arr[i] == 0 {
			// Edge case: If the zero is at the boundary, we only have space to duplicate it once.
			if i == lastIndex-zerosToDuplicate {
				arr[n-1] = 0
				lastIndex--
				break
			}
			zerosToDuplicate++
		}
	}
	// Second Pass: Move elements backward
	// We start from the last valid element and the end of the array.
	source := lastIndex - zerosToDuplicate
	dest := lastIndex

	for source >= 0 {
		if arr[source] == 0 {
			arr[dest] = 0
			dest--
			arr[dest] = 0
		} else {
			arr[dest] = arr[source]
		}
		source--
		dest--
	}
}
