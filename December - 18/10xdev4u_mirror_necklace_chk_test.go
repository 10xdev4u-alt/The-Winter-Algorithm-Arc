package main

import "testing"

func TestChkMirror(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		expected bool
	}{
		{
			name:     "Mirrored Odd Length",
			arr:      []int{1, 2, 3, 2, 1},
			expected: true,
		},
		{
			name:     "Mirrored Even Length",
			arr:      []int{1, 2, 2, 1},
			expected: true,
		},
		{
			name:     "Not Mirrored",
			arr:      []int{1, 2, 3, 4, 5},
			expected: false,
		},
		{
			name:     "Not Mirrored (Close Match)",
			arr:      []int{1, 2, 3, 1}, // 3 != 2
			expected: false,
		},
		{
			name:     "Single Bead",
			arr:      []int{99},
			expected: true,
		},
		{
			name:     "Empty Necklace",
			arr:      []int{},
			expected: true,
		},
		{
			name:     "Negative Numbers (Mirrored)",
			arr:      []int{-10, 5, -10},
			expected: true,
		},
		{
			name:     "Negative Numbers (Not Mirrored)",
			arr:      []int{-10, 5, 10},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Your function requires passing the length explicitly
			result := chkMirror(tt.arr, len(tt.arr))

			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
