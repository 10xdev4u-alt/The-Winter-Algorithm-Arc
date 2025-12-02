package main

import (
	"testing"
)

func TestMinSweetx(t *testing.T) {
	tests := []struct {
		name     string
		scores   []int
		expected int
	}{
		{
			name:     "Sample 1: Valley Pattern",
			scores:   []int{1, 0, 2},
			expected: 5, // Distribution: [2, 1, 2]
		},
		{
			name:     "Sample 2: Plateau",
			scores:   []int{1, 2, 2},
			expected: 4, // Distribution: [1, 2, 1] (Third guy isn't > second, so gets 1)
		},
		{
			name:     "Strictly Increasing",
			scores:   []int{1, 3, 5, 7},
			expected: 10, // [1, 2, 3, 4] -> Sum 10
		},
		{
			name:     "Strictly Decreasing",
			scores:   []int{10, 5, 2},
			expected: 6, // [3, 2, 1] -> Sum 6
		},
		{
			name:     "Mountain Peak",
			scores:   []int{1, 5, 2},
			expected: 4, // [1, 2, 1] -> Sum 4
		},
		{
			name:     "Flat Land",
			scores:   []int{1, 1, 1, 1},
			expected: 4, // [1, 1, 1, 1] -> Sum 4
		},
		{
			name:     "Complex Zig-Zag",
			scores:   []int{1, 3, 2, 2, 1},
			expected: 7, // [1, 2, 1, 2, 1] -> Sum 7
		},
		{
			name:     "Empty Class",
			scores:   []int{},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minSweetx(tt.scores, len(tt.scores))
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}
