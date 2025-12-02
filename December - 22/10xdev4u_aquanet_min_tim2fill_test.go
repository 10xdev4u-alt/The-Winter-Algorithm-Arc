package main

import (
	"testing"
)

func TestMinTime2Fill(t *testing.T) {
	tests := []struct {
		name     string
		v        int
		edges    [][]int
		water    []int
		expected int
	}{
		{
			name:     "Sample 1: The Fork (Correct Input)",
			v:        7,
			edges:    [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {1, 5}, {5, 6}},
			water:    []int{1, 0, 0, 0, 0, 0, 0},
			expected: 4,
			// Path: 0->1->2->3->4 (Distance 4)
		},
		{
			name:     "The Straight Line (Your Typo Input)",
			v:        7,
			edges:    [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}},
			water:    []int{1, 0, 0, 0, 0, 0, 0},
			expected: 6,
			// Path: 0->1->2->3->4->5->6 (Distance 6)
		},
		{
			name:     "Disconnected Graph (Impossible)",
			v:        6,
			edges:    [][]int{{0, 1}, {1, 2}, {3, 4}, {4, 5}},
			water:    []int{1, 1, 0, 0, 0, 0},
			expected: -1,
		},
		{
			name:     "Immediate Fill (All Wet)",
			v:        3,
			edges:    [][]int{{0, 1}, {1, 2}},
			water:    []int{1, 1, 1},
			expected: 0,
		},
		{
			name:     "Two Sources Meeting",
			v:        5,
			edges:    [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}},
			water:    []int{1, 0, 0, 0, 1}, // Sources at 0 and 4
			expected: 2,                    // They meet at node 2
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minTime2Fill(tt.v, tt.edges, tt.water)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}
