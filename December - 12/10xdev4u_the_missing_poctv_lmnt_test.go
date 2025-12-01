package main

import (
	"testing"
)

func TestChkDupndMisng(t *testing.T) {
	tests := []struct {
		name        string
		input       []int
		wantMissing int
		wantDup     int
	}{
		{
			name:        "Example 1 (N=5)",
			input:       []int{3, 1, 2, 2, 5},
			wantMissing: 4,
			wantDup:     2,
		},
		{
			name:        "Example 2 (N=4)",
			input:       []int{4, 3, 3, 1},
			wantMissing: 2,
			wantDup:     3,
		},
		{
			name:        "Duplicate at Start (N=2)",
			input:       []int{1, 1},
			wantMissing: 2,
			wantDup:     1,
		},
		{
			name:        "Duplicate at End (N=2)",
			input:       []int{2, 2},
			wantMissing: 1,
			wantDup:     2,
		},
		{
			name:        "Larger Case",
			input:       []int{1, 2, 3, 4, 4},
			wantMissing: 5,
			wantDup:     4,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Create a copy because your function modifies the array
			inputCopy := make([]int, len(tc.input))
			copy(inputCopy, tc.input)

			gotMissing, gotDup := chkDupndMisng(inputCopy)

			if gotMissing != tc.wantMissing || gotDup != tc.wantDup {
				t.Errorf("\nInput: %v\nGot:  Missing=%d, Dup=%d\nWant: Missing=%d, Dup=%d",
					tc.input, gotMissing, gotDup, tc.wantMissing, tc.wantDup)
			}
		})
	}
}
