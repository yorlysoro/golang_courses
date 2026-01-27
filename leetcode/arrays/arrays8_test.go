package main

import (
	"reflect"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		expect []int
	}{
		{
			name:   "Example 1: Mixed numbers",
			input:  []int{-4, -1, 0, 3, 10},
			expect: []int{0, 1, 9, 16, 100},
		},
		{
			name:   "Example 2: Negative and odd length",
			input:  []int{-7, -3, 2, 3, 11},
			expect: []int{4, 9, 9, 49, 121},
		},
		{
			name:   "All negative numbers",
			input:  []int{-5, -3, -2, -1},
			expect: []int{1, 4, 9, 25},
		},
		{
			name:   "Single element",
			input:  []int{-2},
			expect: []int{4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sortedSquares(tt.input)
			if !reflect.DeepEqual(result, tt.expect) {
				t.Errorf("sortedSquares(%v) = %v; want %v", tt.input, result, tt.expect)
			}
		})
	}
}
