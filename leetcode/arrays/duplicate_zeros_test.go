package main

import (
	"reflect"
	"testing"
)

func TestDuplicateZeros(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Example 1",
			input:    []int{1, 0, 2, 3, 0, 4, 5, 0},
			expected: []int{1, 0, 0, 2, 3, 0, 0, 4},
		},
		{
			name:     "Example 2",
			input:    []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "All zeros",
			input:    []int{0, 0, 0},
			expected: []int{0, 0, 0},
		},
		{
			name:     "Edge case: Zero at the end",
			input:    []int{1, 2, 0},
			expected: []int{1, 2, 0},
		},
		{
			name:     "Edge case: Zero that can't be duplicated due to space",
			input:    []int{8, 4, 5, 0, 0, 0, 0, 7},
			expected: []int{8, 4, 5, 0, 0, 0, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Work on a copy to keep the original test case clean
			arr := make([]int, len(tt.input))
			copy(arr, tt.input)
			duplicateZeros(arr)
			if !reflect.DeepEqual(arr, tt.expected) {
				t.Errorf("duplicateZeros(%v) = %v; want %v", tt.input, arr, tt.expected)
			}
		})
	}
}
