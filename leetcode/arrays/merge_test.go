package main

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name     string
		nums1    []int
		m        int
		nums2    []int
		n        int
		expected []int
	}{
		{
			name:     "Example 1: Standard merge",
			nums1:    []int{1, 2, 3, 0, 0, 0},
			m:        3,
			nums2:    []int{2, 5, 6},
			n:        3,
			expected: []int{1, 2, 2, 3, 5, 6},
		},
		{
			name:     "Example 2: nums2 empty",
			nums1:    []int{1},
			m:        1,
			nums2:    []int{},
			n:        0,
			expected: []int{1},
		},
		{
			name:     "Example 3: nums1 empty (m=0)",
			nums1:    []int{0},
			m:        0,
			nums2:    []int{1},
			n:        1,
			expected: []int{1},
		},
		{
			name:     "Constraints: Negative values",
			nums1:    []int{-5, -1, 0, 0, 0},
			m:        2,
			nums2:    []int{-10, 2, 4},
			n:        3,
			expected: []int{-10, -5, -1, 2, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a local copy to modify in place
			input1 := make([]int, len(tt.nums1))
			copy(input1, tt.nums1)
			merge(input1, tt.m, tt.nums2, tt.n)
			if !reflect.DeepEqual(input1, tt.expected) {
				t.Errorf("merge(%v, %d, %v, %d) = %v; want %v", tt.nums1, tt.m, tt.nums2, tt.n, input1, tt.expected)
			}
		})
	}
}
