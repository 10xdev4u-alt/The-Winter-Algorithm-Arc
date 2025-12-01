package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func captureOutput(f func()) string {
	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f() // Run the function passed in

	w.Close()
	os.Stdout = old // restore the real stdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func TestFormatNo(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		width    int // We calculate width manually for the test cases
		expected string
	}{
		{
			name:  "Example 1 (N=3)",
			n:     3,
			width: 2, // Binary of 3 is "11" -> len 2
			expected: " 1  1  1  1\n" +
				" 2  2  2 10\n" +
				" 3  3  3 11\n",
		},
		{
			name:  "Example 2 (N=6)",
			n:     6,
			width: 3, // Binary of 6 is "110" -> len 3
			expected: "  1   1   1   1\n" +
				"  2   2   2  10\n" +
				"  3   3   3  11\n" +
				"  4   4   4 100\n" +
				"  5   5   5 101\n" +
				"  6   6   6 110\n",
		},
		{
			name:     "Smallest (N=1)",
			n:        1,
			width:    1, // Binary of 1 is "1" -> len 1
			expected: "1 1 1 1\n",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := captureOutput(func() {
				formatNo(tc.n, tc.width)
			})

			if actual != tc.expected {
				t.Errorf("\n--- FAILED: %s ---\nExpected:\n%q\nActual:\n%q", tc.name, tc.expected, actual)
			}
		})
	}
}
