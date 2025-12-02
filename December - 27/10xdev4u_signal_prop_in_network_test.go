package main

import (
	"testing"
)

func TestSignalPropg(t *testing.T) {
	tests := []struct {
		name     string
		N        int
		edges    [][3]int
		k        int // Source node
		expected int
	}{
		{
			name: "Sample 1: Standard Network",
			N:    4,
			edges: [][3]int{
				{0, 1, 2},
				{1, 2, 1},
				{0, 2, 4},
				{2, 3, 3},
			},
			k:        0,
			expected: 6, // 0->1(2)->2(3)->3(6)
		},
		{
			name: "Sample 2: Unreachable Node",
			N:    3,
			edges: [][3]int{
				{0, 1, 5},
				{1, 0, 5},
			},
			k:        0,
			expected: -1, // Node 2 is unreachable
		},
		{
			name: "Alternative Path (Direct is Slower)",
			N:    3,
			edges: [][3]int{
				{0, 1, 10},
				{0, 2, 2},
				{2, 1, 3}, // 0->2->1 = 5
			},
			k:        0,
			expected: 5,
		},
		{
			name:     "Single Node",
			N:        1,
			edges:    [][3]int{},
			k:        0,
			expected: 0,
		},
		{
			name: "Cyclic Graph",
			N:    3,
			edges: [][3]int{
				{0, 1, 1},
				{1, 2, 1},
				{2, 0, 1}, // Cycle
			},
			k:        0,
			expected: 2, // Reaches last node in 2 steps
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := signalPropg(tt.N, tt.edges, tt.k)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}
