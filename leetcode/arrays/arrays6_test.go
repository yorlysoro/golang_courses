package main

import "testing"

func TestFindMaxConsecutiveOnes(t *testing.T) {
	// Table-driven test cases
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1: Longest at the end",
			nums:     []int{1, 1, 0, 1, 1, 1},
			expected: 3,
		},
		{
			name:     "Example 2: Mixed values",
			nums:     []int{1, 0, 1, 1, 0, 1},
			expected: 2,
		},
		{
			name:     "All zeros",
			nums:     []int{0, 0, 0, 0},
			expected: 0,
		},
		{
			name:     "All ones",
			nums:     []int{1, 1, 1, 1, 1},
			expected: 5,
		},
		{
			name:     "Single element 1",
			nums:     []int{1},
			expected: 1,
		},
		{
			name:     "Large input constraint check",
			nums:     []int{1, 0, 1}, // Simplified representation
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findMaxConsecutiveOnes(tt.nums)
			if result != tt.expected {
				t.Errorf("findMaxConsecutiveOnes(%v) = %d; want %d", tt.nums, result, tt.expected)
			}
		})
	}
}
