package main

import (
	"testing"
)

func TestMinWeightCycle(t *testing.T) {
	tests := []struct {
		name     string
		V        int
		edges    [][3]int
		expected int
	}{
		{
			name: "Sample 1: Mixed Graph",
			V:    5,
			edges: [][3]int{
				{0, 1, 2}, {1, 2, 2}, {1, 3, 1},
				{1, 4, 1}, {0, 4, 3}, {2, 3, 4},
			},
			expected: 6, // Cycle: 0-1-4-0 (2+1+3)
		},
		{
			name: "Sample 2: Larger Graph",
			V:    6,
			edges: [][3]int{
				{0, 1, 4}, {1, 2, 3}, {2, 0, 5},
				{1, 3, 2}, {3, 4, 6}, {4, 1, 1},
			},
			expected: 9, // Cycle: 1-3-4-1 (2+6+1)
		},
		{
			name: "Small Triangle",
			V:    3,
			edges: [][3]int{
				{0, 1, 1}, {1, 2, 1}, {2, 0, 1},
			},
			expected: 3,
		},
		{
			name: "No Cycle (Tree)",
			V:    4,
			edges: [][3]int{
				{0, 1, 1}, {1, 2, 1}, {2, 3, 1},
			},
			expected: -1,
		},
		{
			name: "Two Separate Cycles",
			V:    5,
			edges: [][3]int{
				{0, 1, 10}, {1, 2, 10}, {2, 0, 10}, // Cycle A: 30
				{3, 4, 2}, // Disconnected Edge
			},
			expected: 30,
		},
		{
			name: "Disconnected Graph with No Cycle",
			V:    5,
			edges: [][3]int{
				{0, 1, 5}, {3, 4, 5},
			},
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minWeightCycle(tt.V, tt.edges)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}
