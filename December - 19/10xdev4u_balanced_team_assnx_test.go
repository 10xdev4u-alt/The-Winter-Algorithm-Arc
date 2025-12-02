package main

import "testing"

func TestMinDiff(t *testing.T) {
	tests := []struct {
		name     string
		skills   []int
		n        int
		expected int
	}{
		{
			name:     "Sample 1: Perfect Balance",
			skills:   []int{3, 1, 4, 2, 2},
			n:        5,
			expected: 0, // Sum 12. Teams {3,1,2} and {4,2} -> 6 vs 6
		},
		{
			name:     "Sample 2: Slight Imbalance",
			skills:   []int{1, 2, 3, 5},
			n:        4,
			expected: 1, // Sum 11. Teams {5,1} and {3,2} -> 6 vs 5
		},
		{
			name:     "Greedy Trap Case",
			skills:   []int{5, 5, 4, 4, 2},
			n:        5,
			expected: 0, // Teams {5,5} and {4,4,2} -> 10 vs 10
		},
		{
			name:     "Two Person Gap",
			skills:   []int{1, 100},
			n:        2,
			expected: 99, // 100 - 1
		},
		{
			name:     "Single Person",
			skills:   []int{10},
			n:        1,
			expected: 10, // Team A {10}, Team B {}
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minDiff(tt.skills, tt.n)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}
