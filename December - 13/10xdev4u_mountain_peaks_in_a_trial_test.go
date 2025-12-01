package main

import (
	"reflect"
	"testing"
)

func TestCntPeaks(t *testing.T) {
	tests := []struct {
		name    string
		heights []int
		n       int
		want    []int
	}{
		{
			name:    "Sample 1: Two Peaks",
			heights: []int{1, 3, 2, 4, 5, 3, 2, 1},
			n:       8,
			want:    []int{1, 4},
		},
		{
			name:    "Sample 2: Increasing (No Peak)",
			heights: []int{1, 2, 3, 4, 5},
			n:       5,
			want:    nil, // No peaks found
		},
		{
			name:    "Decreasing (No Peak)",
			heights: []int{5, 4, 3, 2, 1},
			n:       5,
			want:    nil,
		},
		{
			name:    "Zig Zag",
			heights: []int{1, 10, 1, 20, 1},
			n:       5,
			want:    []int{1, 3},
		},
		{
			name:    "Plateau (Strictly Greater Check)",
			heights: []int{2, 5, 5, 2},
			n:       4,
			want:    nil, // 5 is not > 5
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := cntPeaks(tc.heights, tc.n)

			// Handle nil vs empty slice comparison
			if len(got) == 0 && len(tc.want) == 0 {
				return // Both correspond to "no peaks"
			}

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("\nInput: %v\nGot:  %v\nWant: %v", tc.heights, got, tc.want)
			}
		})
	}
}
