package main

import "testing"

func TestFindNumbers(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1: Mixed digits",
			nums:     []int{12, 345, 2, 6, 7896},
			expected: 2,
		},
		{
			name:     "Example 2: Only one even-length number",
			nums:     []int{555, 901, 482, 1771},
			expected: 1,
		},
		{
			name:     "Constraint Check: Max value 10^5 (6 digits)",
			nums:     []int{100000, 1},
			expected: 1, // 100,000 has 6 digits (even)
		},
		{
			name:     "Single digit numbers",
			nums:     []int{1, 5, 9},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := findNumbers(tt.nums)
			if actual != tt.expected {
				t.Errorf("%s: expected %d, got %d", tt.name, tt.expected, actual)
			}
		})
	}
}
