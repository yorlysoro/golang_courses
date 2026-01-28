package main

// merge combines nums2 into nums1 as a single sorted array.
// It modifies nums1 in-place.

func merge(nums1 []int, m int, nums2 []int, n int) {
	// Pointers for the end of valid elements in nums1, nums2, and the total capacity
	i, j, k := m-1, n-1, m+n-1

	// While there are elements in both arrays to compare
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}

	// If there are remaining elements in nums2, copy them.
	// (Note: elements remaining in nums1 are already in their correct place)
	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}

/*
Complexity Analysis:
Time Complexity: O(m+n). We visit each element at most once.

Space Complexity: O(1). We are not using any extra arrays or data structures; we use the space already provided in nums1.

Cleanliness: By using a simple for loop and a single cleanup loop for nums2, the code remains readable and efficient.
*/
