package main

import (
	"testing"
)

func TestMagicalChestParser(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string // We expect the output to match the input exactly
	}{
		{
			name:  "Simple List",
			input: "[1,2,3,4,5]",
			want:  "[1,2,3,4,5]",
		},
		{
			name:  "Single Nesting",
			input: "[1,[2,3],4]",
			want:  "[1,[2,3],4]",
		},
		{
			name:  "Double Nesting",
			input: "[1,[4,[6]]]",
			want:  "[1,[4,[6]]]",
		},
		{
			name:  "Empty List",
			input: "[]",
			want:  "[]",
		},
		{
			name:  "Empty Nested Lists",
			input: "[[],[[]],[]]",
			want:  "[[],[[]],[]]",
		},
		{
			name:  "Negative Numbers",
			input: "[-10,2,-500]",
			want:  "[-10,2,-500]",
		},
		{
			name:  "Multi-digit Numbers",
			input: "[123,[456,789]]",
			want:  "[123,[456,789]]",
		},
		{
			name:  "Complex Mix",
			input: "[1,[2,[3,[4,[5,6,7]]]],8,9]",
			want:  "[1,[2,[3,[4,[5,6,7]]]],8,9]",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// 1. Parse the string into the struct
			parsedStruct := parse(tc.input)

			// 2. Convert the struct back to a string
			got := parsedStruct.String()

			// 3. Compare
			if got != tc.want {
				t.Errorf("\nInput: %s\nGot:   %s\nWant:  %s", tc.input, got, tc.want)
			}
		})
	}
}
