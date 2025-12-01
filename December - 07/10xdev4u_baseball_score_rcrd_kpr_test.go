package main

import "testing"

func TestCalPoints(t *testing.T) {
	tests := []struct {
		name     string
		ops      []string
		expected int
	}{
		{
			name:     "Example 1: Mixed Operations",
			ops:      []string{"5", "2", "C", "D", "+"},
			expected: 30,
		},
		{
			name:     "Example 2: Negative Numbers",
			ops:      []string{"5", "-2", "4", "C", "D", "9", "+", "+"},
			expected: 27,
		},
		{
			name:     "Simple Addition",
			ops:      []string{"1", "1", "+"}, // [1, 1, 2] -> 4
			expected: 4,
		},
		{
			name:     "Only Numbers",
			ops:      []string{"10", "20", "30"},
			expected: 60,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calPoints(tt.ops)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}
