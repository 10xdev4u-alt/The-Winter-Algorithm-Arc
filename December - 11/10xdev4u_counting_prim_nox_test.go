package main

import (
	"testing"
)

func TestCntPrime(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "Example 1 (N=10)",
			n:    10,
			want: 4, // 2, 3, 5, 7
		},
		{
			name: "Example 2 (N=20)",
			n:    20,
			want: 8, // 2, 3, 5, 7, 11, 13, 17, 19
		},
		{
			name: "Edge Case (N=0)",
			n:    0,
			want: 0,
		},
		{
			name: "Edge Case (N=2)",
			n:    2,
			want: 0, // Strictly less than 2
		},
		{
			name: "Small Prime (N=3)",
			n:    3,
			want: 1, // Just 2
		},
		{
			name: "Large Number (N=100)",
			n:    100,
			want: 25,
		},
		{
			name: "Stress Test (N=1000)",
			n:    1000,
			want: 168,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := cntPrime(tc.n)
			if got != tc.want {
				t.Errorf("Input N=%d\nGot: %d\nWant: %d", tc.n, got, tc.want)
			}
		})
	}
}
