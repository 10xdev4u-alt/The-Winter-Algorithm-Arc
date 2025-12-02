package main

import "testing"

func TestFistUniqChar(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Sample 1: Mixed",
			input:    "swiss",
			expected: "w", // 's' repeats, 'w' is first unique
		},
		{
			name:     "Sample 2: No Unique",
			input:    "aabbcc",
			expected: "-1",
		},
		{
			name:     "Unique at End",
			input:    "loveleetcode",
			expected: "v", // 'l', 'o', 'e' repeat. 'v' is unique.
		},
		{
			name:     "Unique at Start",
			input:    "z",
			expected: "z",
		},
		{
			name:     "All Unique",
			input:    "abcde",
			expected: "a",
		},
		{
			name:     "Long Pattern",
			input:    "goforgood",
			expected: "f", // 'g', 'o' repeat. 'f' is first.
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := fistUniqChar(tt.input)
			if result != tt.expected {
				t.Errorf("got %q, want %q", result, tt.expected)
			}
		})
	}
}
