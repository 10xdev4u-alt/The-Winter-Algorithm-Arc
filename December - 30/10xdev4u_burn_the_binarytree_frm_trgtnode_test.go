package main

import (
	"testing"
)

func TestBurnTree(t *testing.T) {
	tests := []struct {
		name     string
		tree     []int // Level-order representation (-1 for null)
		target   int
		expected [][]int
	}{
		{
			name: "Sample 1: Target 14",
			// Tree: [12, 13, 10, -1, -1, 14, 24, -1, -1, -1, -1, 21, -1, -1, -1]
			// Simplified Level Order for Sample 1 structure
			tree: []int{12, 13, 10, -1, -1, 14, 24, 21},
			// Note: My buildTree helper logic assigns children sequentially.
			// 12 -> L:13, R:10
			// 13 -> L:nil, R:nil
			// 10 -> L:14, R:24
			// 14 -> L:21, R:nil
			// This matches Sample 1 structure.
			target: 14,
			expected: [][]int{
				{14},
				{21, 10}, // Left child, Parent
				{12, 24}, // Parent's Parent, Parent's Right Child
				{13},     // Parent's Parent's Left Child
			},
			// Note: The order inside a level (e.g., {21, 10} vs {10, 21}) depends on
			// exact implementation (Left/Right/Parent check order).
			// The test helper below handles unordered comparisons for inner slices.
		},
		{
			name: "Sample 2: Target 41 (Root)",
			tree: []int{41, 2, 19, 12, -1, -1, -1},
			// 41 -> 2, 19
			// 2 -> 12
			target: 41,
			expected: [][]int{
				{41},
				{2, 19},
				{12},
			},
		},
		{
			name:   "Leaf Node Target",
			tree:   []int{1, 2, 3},
			target: 2,
			expected: [][]int{
				{2},
				{1},
				{3},
			},
		},
		{
			name:   "Linear Tree (Line)",
			tree:   []int{1, 2, -1, 3, -1, 4}, // 1->2->3->4
			target: 3,
			expected: [][]int{
				{3},
				{4, 2},
				{1},
			},
		},
		{
			name:     "Target Not Found",
			tree:     []int{1, 2, 3},
			target:   99,
			expected: [][]int{},
		},
		{
			name:   "Single Node",
			tree:   []int{1},
			target: 1,
			expected: [][]int{
				{1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTreeFromLevelOrder(tt.tree)
			result := burnTree(root, tt.target)

			if len(result) != len(tt.expected) {
				t.Fatalf("Length mismatch: got %d levels, want %d", len(result), len(tt.expected))
			}

			for i := 0; i < len(result); i++ {
				if !elementsMatch(result[i], tt.expected[i]) {
					t.Errorf("Level %d mismatch: got %v, want %v", i, result[i], tt.expected[i])
				}
			}
		})
	}
}

// Helper to compare slices ignoring order (sets)
func elementsMatch(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	counts := make(map[int]int)
	for _, x := range a {
		counts[x]++
	}
	for _, x := range b {
		counts[x]--
		if counts[x] < 0 {
			return false
		}
	}
	return true
}
