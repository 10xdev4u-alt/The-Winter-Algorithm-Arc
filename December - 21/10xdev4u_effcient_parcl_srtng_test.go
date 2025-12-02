package main

import (
	"reflect"
	"testing"
)

func TestQueueSrt(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Sample 1",
			input:    []int{4, 2, 1, 5, 3},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Sample 2 (Tens)",
			input:    []int{10, 30, 20, 40},
			expected: []int{10, 20, 30, 40},
		},
		{
			name:     "Reverse Sorted",
			input:    []int{6, 5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "Already Sorted",
			input:    []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Duplicates",
			input:    []int{5, 1, 5, 1},
			expected: []int{1, 1, 5, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// We pass a copy because your function modifies the input slice
			inputCopy := make([]int, len(tt.input))
			copy(inputCopy, tt.input)

			result := queueSrt(inputCopy)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
