package main

import (
	"reflect"
	"testing"
)

// Helper: Convert Linked List back to Slice for easy comparison
func listToSlice(head *Node) []int {
	var result []int
	curr := head
	for curr != nil {
		result = append(result, curr.data)
		curr = curr.next
	}
	return result
}

func TestReOrderLL(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Even Length (N=4)",
			input:    []int{1, 2, 3, 4},
			expected: []int{1, 4, 2, 3},
		},
		{
			name:     "Odd Length (N=5)",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 5, 2, 4, 3},
		},
		{
			name:     "Small List (N=2)",
			input:    []int{10, 20},
			expected: []int{10, 20},
		},
		{
			name:     "Longer Even (N=6)",
			input:    []int{1, 2, 3, 4, 5, 6},
			expected: []int{1, 6, 2, 5, 3, 4},
		},
		{
			name:     "Single Node",
			input:    []int{99},
			expected: []int{99},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// 1. Build the list
			head := createLL(tc.input)

			// 2. Run your logic
			reOrderLL(head)

			// 3. Convert back to slice to check
			got := listToSlice(head)

			// 4. Compare
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("\nInput:    %v\nExpected: %v\nGot:      %v", tc.input, tc.expected, got)
			}
		})
	}
}
