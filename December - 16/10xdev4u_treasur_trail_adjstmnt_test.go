package main

import (
	"reflect"
	"testing"
)

// Helper to convert your linked list back to a slice for comparison
func listToSlice(head *node) []int {
	var result []int
	curr := head
	for curr != nil {
		result = append(result, curr.val)
		curr = curr.next
	}
	return result
}

func TestRemovNth(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		n     int
		want  []int
	}{
		{
			name:  "Example 1: Remove 2nd from end",
			input: []int{1, 2, 3, 4, 5},
			n:     2,
			want:  []int{1, 2, 3, 5},
		},
		{
			name:  "Example 2: Single Node (Become Empty)",
			input: []int{1},
			n:     1,
			want:  nil, // Expecting nil/empty
		},
		{
			name:  "Example 3: Remove Last Node",
			input: []int{1, 2},
			n:     1,
			want:  []int{1},
		},
		{
			name:  "Remove First Node (Head)",
			input: []int{10, 20, 30},
			n:     3, // 3rd from end is 10
			want:  []int{20, 30},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// 1. Create List using your function
			// Note: Your createLL assumes array is not empty, which matches constraints
			head := createLL(tc.input, len(tc.input))

			// 2. Run logic
			newHead := removNth(head, tc.n)

			// 3. Convert result back to slice
			got := listToSlice(newHead)

			// 4. Compare
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("\nInput: %v (n=%d)\nGot:  %v\nWant: %v", tc.input, tc.n, got, tc.want)
			}
		})
	}
}
