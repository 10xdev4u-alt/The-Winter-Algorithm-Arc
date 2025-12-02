package main

import (
	"testing"
)

func TestLongestPth(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name: "Sample 1: Mixed Path",
			grid: [][]int{
				{9, 9, 4},
				{6, 6, 8},
				{2, 1, 1},
			},
			expected: 4, // 1->2->6->9
		},
		{
			name: "Sample 2: Snake Path",
			grid: [][]int{
				{3, 4, 5},
				{3, 2, 6},
				{2, 2, 1},
			},
			expected: 4, // 3->4->5->6
		},
		{
			name: "Single Cell",
			grid: [][]int{
				{1},
			},
			expected: 1,
		},
		{
			name: "Strictly Increasing Row",
			grid: [][]int{
				{1, 2, 3, 4, 5},
			},
			expected: 5,
		},
		{
			name: "Strictly Decreasing (No steps)",
			grid: [][]int{
				{5, 4, 3, 2, 1},
			},
			expected: 5, // The path starts at 1 and goes to 5 (or 5 to 1 depends on logic, increasing is 1->5)
			// Wait, strictly ASCENDING path.
			// If we start at 1, neighbors are 2. 1->2->3->4->5. Length 5.
		},
		{
			name: "Pyramid (Peak in middle)",
			grid: [][]int{
				{1, 2, 1},
				{2, 5, 2},
				{1, 2, 1},
			},
			expected: 3, // 1->2->5
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := len(tt.grid)
			c := len(tt.grid[0])

			// Your function signature
			result := longestPth(tt.grid, r, c)

			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}
