package main

import (
	"testing"
)

func TestTreasureHunt(t *testing.T) {
	tests := []struct {
		name     string
		grid     []string
		expected int
	}{
		{
			name: "Sample 1: Basic Key and Door",
			grid: []string{
				"S.a",
				"#A#",
				"..T",
			},
			expected: 6,
			// Path: S->(0,1)->a(0,2)[get key]->(0,1)->(1,1)[open A]->(2,1)->T(2,2)
		},
		{
			name: "Sample 2: Unreachable (Blocked by Wall)",
			grid: []string{
				"S.A.",
				"###.",
				"a...",
			},
			expected: -1,
		},
		{
			name: "Direct Path (No Keys Needed)",
			grid: []string{
				"S..",
				"...",
				"..T",
			},
			expected: 4,
		},
		{
			name: "Backtracking Required (Key behind Start)",
			grid: []string{
				"a.S.A.T",
			},
			expected: 8,
			// S(0,2) -> a(0,0) [steps=2] -> S(0,2) [steps=4] -> A(0,4) [steps=6] -> T(0,6) [steps=8]
		},
		{
			name: "Multiple Keys (a and b needed)",
			grid: []string{
				"S.a.b",
				"#####",
				".B.A.",
				"#####",
				"..T..",
			},
			expected: -1, // Impossible because of wall
		},
		{
			name: "Multi Key Success",
			grid: []string{
				"S.a",
				"...",
				"b.A",
				"...",
				"B.T",
			},
			expected: 6,
			// S->a (get A), down to b (get B), up to A (open), down to B (open), to T.
		},
		{
			name: "Door Without Key (Fail)",
			grid: []string{
				"S.A.T",
			},
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rows := len(tt.grid)
			cols := len(tt.grid[0])

			// Your function signature
			result := tresurHnt(tt.grid, rows, cols)

			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}
