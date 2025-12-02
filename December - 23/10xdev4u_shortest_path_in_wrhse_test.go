package main

import (
	"testing"
)

func TestShortsPth(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name: "Sample 1: Standard Path",
			grid: [][]int{
				{0, 0, 0, 0, 0},
				{1, 1, 0, 1, 0},
				{0, 0, 0, 0, 0},
				{0, 1, 1, 1, 0},
			},
			expected: 7,
		},
		{
			name: "Sample 2: U-Turn Path",
			grid: [][]int{
				{0, 1, 0},
				{0, 1, 0},
				{0, 0, 0},
			},
			expected: 4,
		},
		{
			name: "Direct Line (All 0s)",
			grid: [][]int{
				{0, 0, 0},
				{0, 0, 0},
			},
			expected: 3, // (0,0)->(0,1)->(0,2)->(1,2)
		},
		{
			name: "Unreachable (Blocked by 1s)",
			grid: [][]int{
				{0, 1},
				{1, 0},
			},
			expected: -1,
		},
		{
			name: "Start Blocked",
			grid: [][]int{
				{1, 0},
				{0, 0},
			},
			expected: -1,
		},
		{
			name: "End Blocked",
			grid: [][]int{
				{0, 0},
				{0, 1},
			},
			expected: -1,
		},
		{
			name: "1x1 Grid (Start is End)",
			grid: [][]int{
				{0},
			},
			expected: 0,
		},
		{
			name: "Complex Snake Maze",
			grid: [][]int{
				{0, 0, 0, 1, 0},
				{1, 1, 0, 1, 0},
				{0, 0, 0, 0, 0},
				{0, 1, 1, 1, 1},
				{0, 0, 0, 0, 0},
			},
			expected: 12, // Needs to wind all the way around
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rows := len(tt.grid)
			cols := len(tt.grid[0])

			// Your function expects (grid, rows, cols)
			result := shortsPth(tt.grid, rows, cols)

			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}
