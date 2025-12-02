package main

import (
	"testing"
)

func TestIsCmplTree(t *testing.T) {
	tests := []struct {
		name     string
		vals     []int // We use your buldTre logic: -1 represents null
		expected bool
	}{
		{
			name:     "Example 1: Complete Tree",
			vals:     []int{1, 2, 3, 4, 5, 6},
			expected: true,
		},
		{
			name:     "Example 2: Gap in Last Level",
			vals:     []int{1, 2, 3, 4, 5, -1, 7}, // 7 comes after a null (-1)
			expected: false,
		},
		{
			name:     "Full Binary Tree",
			vals:     []int{1, 2, 3},
			expected: true,
		},
		{
			name:     "Missing Left Child (Bad)",
			vals:     []int{1, -1, 2},
			expected: false,
		},
		{
			name:     "Missing Right Child (Good - Left Filled)",
			vals:     []int{1, 2},
			expected: true,
		},
		{
			name:     "Empty Tree",
			vals:     []int{},
			expected: true,
		},
		{
			name:     "Single Node",
			vals:     []int{100},
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// 1. Build the tree using your helper
			root := buldTre(tc.vals)

			// 2. Check Logic
			got := isCmplTree(root)

			if got != tc.expected {
				t.Errorf("Tree Input: %v\nGot: %v\nWant: %v", tc.vals, got, tc.expected)
			}
		})
	}
}
