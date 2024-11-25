package main

import (
	"testing"
)

func TestMaxPathSum(t *testing.T) {
	tests := []struct {
		name     string
		triangle [][]int
		expected int
	}{
		{
			name: "Small Triangle",
			triangle: [][]int{
				{59},
				{73, 41},
				{52, 40, 53},
				{26, 53, 6, 34},
			},
			expected: 237,
		},
		{
			name: "Simple Triangle",
			triangle: [][]int{
				{1},
				{2, 3},
				{4, 5, 6},
			},
			expected: 10, // Path: 1 -> 3 -> 6
		},
		{
			name: "Single Node",
			triangle: [][]int{
				{42},
			},
			expected: 42,
		},
		{
			name:     "Empty Triangle",
			triangle: [][]int{},
			expected: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := maxPathSum(test.triangle)
			if result != test.expected {
				t.Errorf("expected %d but got %d", test.expected, result)
			}
		})
	}
}
