package main

import (
	"testing"
)

func TestCountPerfectSquares(t *testing.T) {
	tests := []struct {
		name     string
		inputN   int
		expected int
	}{
		{"Standard Case (N=20)", 20, 4}, // 1, 4, 9, 16
		{"Small Case (N=5)", 5, 2},      // 1, 4

		{
			name:     "The Singularity (N=1)",
			inputN:   1,
			expected: 1, // 1^2 = 1
		},
		{
			name:     "The Cliff Hanger (N=24)",
			inputN:   24,
			expected: 4, // Stops at 16
		},
		{
			name:     "The Exact Hit (N=25)",
			inputN:   25,
			expected: 5, // Includes 25
		},
		{
			name:     "Maximum Constraint (N=100,000)",
			inputN:   100000,
			expected: 316, // sqrt(100,000) approx 316.22
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countPerfectSquares(tt.inputN)

			if got != tt.expected {
				t.Errorf("❌ %s Failed: For N=%d, expected %d but got %d",
					tt.name, tt.inputN, tt.expected, got)
			}
		})
	}
}

func Benchmark_Standard(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countPerfectSquares(1000)
	}
}

func Benchmark_MaxLoad(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countPerfectSquares(100000)
	}
}
