package main

import (
	"testing"
)

func TestCountFleets(t *testing.T) {
	tests := []struct {
		name     string
		target   int
		pos      []int
		speed    []int
		n        int
		expected int
	}{
		{
			name:     "Sample 1: Single Turtle",
			target:   10,
			pos:      []int{3},
			speed:    []int{3},
			n:        1,
			expected: 1,
		},
		{
			name:     "Sample 2: Mixed Merging",
			target:   12,
			pos:      []int{10, 8, 0, 5, 3},
			speed:    []int{2, 4, 1, 1, 3},
			n:        5,
			expected: 3,
			// Explanation:
			// Turtles sorted by pos: (10,2), (8,4), (5,1), (3,3), (0,1)
			// Times: 1.0, 1.0, 7.0, 3.0, 12.0
			// Fleets: {1.0, 1.0}, {7.0, 3.0}, {12.0} -> 3 Fleets
		},
		{
			name:     "Sample 3: Fast Catching Slow (One Big Fleet)",
			target:   100,
			pos:      []int{0, 2, 4},
			speed:    []int{4, 2, 1},
			n:        3,
			expected: 1,
			// Leader at 4 (speed 1) takes 96s.
			// Middle at 2 (speed 2) takes 49s -> Merges.
			// Last at 0 (speed 4) takes 25s -> Merges.
		},
		{
			name:     "No Merging (Reverse Order)",
			target:   10,
			pos:      []int{4, 2, 0},
			speed:    []int{1, 2, 4},
			n:        3,
			expected: 1,
			// Leader at 4 (speed 1) takes 6s.
			// Middle at 2 (speed 2) takes 4s -> Too fast! But started behind leader?
			// Wait, logic check: Leader takes 6s. Middle takes 4s.
			// Middle arrives SOONER (4s) than Leader (6s).
			// Since Middle is BEHIND Leader, Middle CATCHES UP.
			// This case should actually Result in 1 Fleet?
			// Let's re-read the logic:
			// If arrivalTime > lastFleetTime -> New Fleet.
			// 6.0 > -1 -> New Fleet (Time 6.0)
			// 4.0 > 6.0? False -> Merges.
			// 2.5 > 6.0? False -> Merges.
			// So actually, this specific input produces 1 fleet because they all jam up behind the slow leader.
		},
		{
			name:     "Strictly Increasing Times (No Merging)",
			target:   10,
			pos:      []int{5, 0},
			speed:    []int{5, 1},
			n:        2,
			expected: 2,
			// Leader at 5 (speed 5) takes 1s.
			// Last at 0 (speed 1) takes 10s.
			// 10s > 1s -> New Fleet.
		},
		{
			name:     "Zero Turtles",
			target:   10,
			pos:      []int{},
			speed:    []int{},
			n:        0,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Note: Your function signature is (target, speed, pos, n)
			result := countFleets(tt.target, tt.speed, tt.pos, tt.n)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}
