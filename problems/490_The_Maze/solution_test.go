package _90_The_Maze

import "testing"

func TestHasPath(t *testing.T) {
	testcases := []struct {
		maze        [][]int
		start       []int
		destination []int
		expected    bool
	}{
		{
			maze: [][]int{
				{0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 1, 0},
				{1, 1, 0, 1, 1},
				{0, 0, 0, 0, 0},
			},
			start:       []int{0, 4},
			destination: []int{4, 4},
			expected:    true,
		},
		{
			maze: [][]int{
				{0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 0, 1, 0},
				{1, 1, 0, 1, 1},
				{0, 0, 0, 0, 0},
			},
			start:       []int{0, 4},
			destination: []int{3, 2},
			expected:    false,
		},
		{
			maze: [][]int{
				{0, 0, 0, 0, 1, 0, 0},
				{0, 0, 1, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 1},
				{0, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{0, 0, 1, 0, 0, 0, 1},
				{0, 0, 0, 0, 1, 0, 9},
			},
			start:       []int{0, 0},
			destination: []int{8, 6},
			expected:    true,
		},
	}

	for _, tc := range testcases {
		result := hasPath(tc.maze, tc.start, tc.destination)
		if result != tc.expected {
			t.Errorf("Expected %v, but got %v", tc.expected, result)
		}
	}
}
