package main

import (
	"testing"
)

func TestSolveSudoku(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]byte
		solvable bool
	}{
		{
			name: "Sample 1: Solvable",
			input: [][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			solvable: true,
		},
		{
			name: "Impossible Puzzle (Conflict)",
			input: [][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'5', '.', '.', '1', '9', '5', '.', '.', '.'}, // 5 repeated in Col 0
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			solvable: false, // Should fail immediately or return false
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a deep copy because the solver modifies the board
			boardCopy := make([][]byte, 9)
			for i := range tt.input {
				boardCopy[i] = make([]byte, 9)
				copy(boardCopy[i], tt.input[i])
			}

			result := solveSudoku(boardCopy)
			if result != tt.solvable {
				t.Errorf("got %v, want %v", result, tt.solvable)
			}
		})
	}
}
