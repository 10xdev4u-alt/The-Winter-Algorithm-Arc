package main

import (
	"testing"
)

func TestCanCrossBrig(t *testing.T) {
	// 1. Define the scenarios
	tests := []struct {
		name   string // Description of the test
		stones []int  // Input array
		want   bool   // Expected result
	}{
		{
			name:   "Example 1: Possible Path",
			stones: []int{2, 3, 1, 1, 4},
			want:   true,
		},
		{
			name:   "Example 2: Stuck at Zero",
			stones: []int{3, 2, 1, 0, 4},
			want:   false,
		},
		{
			name:   "Single Stone (Start is End)",
			stones: []int{0},
			want:   true, // You are already at the last stone
		},
		{
			name:   "Immediate Dead End",
			stones: []int{0, 2, 3},
			want:   false, // Cannot move from index 0
		},
		{
			name:   "Exact Reach",
			stones: []int{2, 0, 0},
			want:   true, // Index 0 jumps 2 spots to Index 2 (end)
		},
		{
			name:   "Just Short",
			stones: []int{1, 0, 0},
			want:   false, // Index 0 jumps to Index 1 (value 0), stuck
		},
		{
			name:   "Massive Jump",
			stones: []int{10, 0, 0, 0, 0},
			want:   true, // Can easily clear the river
		},
	}

	// 2. Loop through scenarios
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := canCrossBrig(tc.stones)
			if got != tc.want {
				t.Errorf("\nInput: %v\nGot:  %v\nWant: %v", tc.stones, got, tc.want)
			}
		})
	}
}
