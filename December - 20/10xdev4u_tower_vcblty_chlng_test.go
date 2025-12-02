package main

import (
	"reflect"
	"testing"
)

func TestNextTallr(t *testing.T) {
	tests := []struct {
		name     string
		heights  []int
		n        int
		expected []int
	}{
		{
			name:     "Sample 1: Standard Case",
			heights:  []int{4, 5, 2, 25, 7, 6},
			n:        6,
			expected: []int{5, 25, 25, -1, -1, -1},
		},
		{
			name:     "Sample 2: Mixed",
			heights:  []int{13, 7, 6, 12, 10},
			n:        5,
			expected: []int{-1, 12, 12, -1, -1},
		},
		{
			name:     "Strictly Increasing",
			heights:  []int{1, 2, 3, 4, 5},
			n:        5,
			expected: []int{2, 3, 4, 5, -1},
		},
		{
			name:     "Strictly Decreasing (No one sees anything)",
			heights:  []int{5, 4, 3, 2, 1},
			n:        5,
			expected: []int{-1, -1, -1, -1, -1},
		},
		{
			name:     "Equal Heights (Strictly Taller Check)",
			heights:  []int{5, 5, 5},
			n:        3,
			expected: []int{-1, -1, -1}, // 5 cannot see another 5
		},
		{
			name:     "Single Tower",
			heights:  []int{100},
			n:        1,
			expected: []int{-1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := nextTallr(tt.heights, tt.n)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
