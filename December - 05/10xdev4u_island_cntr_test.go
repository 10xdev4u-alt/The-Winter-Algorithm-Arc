package main

import (
	"testing"
)

func TestCountIslands(t *testing.T) {
	tests := []struct {
		name string
		rows int
		cols int
		grid [][]int
		want int
	}{
		{
			name: "Example 1: Two Separate Islands",
			rows: 4, cols: 5,
			grid: [][]int{
				{1, 1, 0, 0, 0},
				{1, 1, 0, 0, 1},
				{0, 0, 0, 1, 1},
				{0, 0, 0, 0, 0},
			},
			want: 2,
		},
		{
			name: "Example 2: Checkerboard (5 Islands)",
			rows: 3, cols: 3,
			grid: [][]int{
				{1, 0, 1},
				{0, 1, 0},
				{1, 0, 1},
			},
			want: 5,
		},
		{
			name: "Single Big Block",
			rows: 2, cols: 2,
			grid: [][]int{
				{1, 1},
				{1, 1},
			},
			want: 1,
		},
		{
			name: "All Water",
			rows: 2, cols: 2,
			grid: [][]int{
				{0, 0},
				{0, 0},
			},
			want: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Make a deep copy of grid because your function modifies it
			gridCopy := make([][]int, tc.rows)
			for i := range tc.grid {
				gridCopy[i] = make([]int, tc.cols)
				copy(gridCopy[i], tc.grid[i])
			}

			got := countIslands(gridCopy, tc.rows, tc.cols)
			if got != tc.want {
				t.Errorf("Got %d, want %d", got, tc.want)
			}
		})
	}
}
