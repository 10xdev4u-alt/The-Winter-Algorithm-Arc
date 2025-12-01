package main

import (
	"testing"
)

func TestMagicSquareLogic(t *testing.T) {
	tests := []struct {
		n    int
		want int // Expected Magic Constant
	}{
		{1, 1},
		{3, 15},
		{5, 65},
		{7, 175},
	}

	for _, tc := range tests {
		// 1. Check Magic Constant
		calculatedMagic := tc.n * (tc.n*tc.n + 1) / 2
		if calculatedMagic != tc.want {
			t.Errorf("For n=%d, expected magic constant %d, got %d", tc.n, tc.want, calculatedMagic)
		}

		// 2. Generate Matrix
		matx := calcMagicMatrx(tc.n)

		// 3. Verify Rows, Cols, and Diagonals
		// Check Rows
		for i := 0; i < tc.n; i++ {
			rowSum := 0
			for j := 0; j < tc.n; j++ {
				rowSum += matx[i][j]
			}
			if rowSum != tc.want {
				t.Errorf("For n=%d, Row %d sum = %d, want %d", tc.n, i, rowSum, tc.want)
			}
		}

		// Check Cols
		for j := 0; j < tc.n; j++ {
			colSum := 0
			for i := 0; i < tc.n; i++ {
				colSum += matx[i][j]
			}
			if colSum != tc.want {
				t.Errorf("For n=%d, Col %d sum = %d, want %d", tc.n, j, colSum, tc.want)
			}
		}

		// Check Main Diagonal
		diagSum := 0
		for i := 0; i < tc.n; i++ {
			diagSum += matx[i][i]
		}
		if diagSum != tc.want {
			t.Errorf("For n=%d, Main Diag sum = %d, want %d", tc.n, diagSum, tc.want)
		}
	}
}
