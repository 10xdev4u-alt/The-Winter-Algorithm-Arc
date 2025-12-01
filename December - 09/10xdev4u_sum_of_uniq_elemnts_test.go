package main

import (
	"testing"
)

func TestSum0fUniq(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Example 1: Mixed",
			input:    []int{1, 2, 2, 3, 4, 4},
			expected: 4, // 1 + 3 = 4
		},
		{
			name:     "Example 2: All Duplicates",
			input:    []int{5, 5, 5, 5, 5},
			expected: 0,
		},
		{
			name:     "All Unique",
			input:    []int{10, 20, 30},
			expected: 60, // 10 + 20 + 30
		},
		{
			name:     "Negative Numbers",
			input:    []int{-5, -5, 2, -1},
			expected: 1, // 2 + (-1) = 1
		},
		{
			name:     "Single Element",
			input:    []int{99},
			expected: 99,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Calling your function 'sum0fUniq'
			got := sum0fUniq(tc.input)
			if got != tc.expected {
				t.Errorf("Input: %v\nGot: %d\nWant: %d", tc.input, got, tc.expected)
			}
		})
	}
}
