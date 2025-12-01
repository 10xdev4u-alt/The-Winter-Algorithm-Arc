package main

import (
	"testing"
)

func TestCountStud(t *testing.T) {
	tests := []struct {
		name     string
		studQ    []int
		sandQ    []int
		expected int
	}{
		{
			name:     "Example 1: Everyone Eats",
			studQ:    []int{1, 1, 0, 0},
			sandQ:    []int{0, 1, 0, 1},
			expected: 0,
		},
		{
			name:     "Example 2: Stuck at Sandwich 0",
			studQ:    []int{1, 1, 1, 0, 0, 1},
			sandQ:    []int{1, 0, 0, 0, 1, 1},
			expected: 3,
		},
		{
			name:     "Immediate Deadlock (Need 1, Top is 0)",
			studQ:    []int{1, 1, 1},
			sandQ:    []int{0, 1, 1},
			expected: 3,
		},
		{
			name:     "Immediate Deadlock (Need 0, Top is 1)",
			studQ:    []int{0, 0, 0},
			sandQ:    []int{1, 0, 0},
			expected: 3,
		},
		{
			name:     "Alternating Perfect Match",
			studQ:    []int{0, 1, 0, 1},
			sandQ:    []int{1, 0, 1, 0},
			expected: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := countStud(tc.studQ, tc.sandQ)
			if got != tc.expected {
				t.Errorf("Input Students: %v, Sandwiches: %v\nGot: %d\nWant: %d",
					tc.studQ, tc.sandQ, got, tc.expected)
			}
		})
	}
}
