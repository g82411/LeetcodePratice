package _73_Sliding_Puzzle

import "testing"

func TestSlidingPuzzle(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		expect := 1
		actual := slidingPuzzle([][]int{{1, 2, 3}, {4, 0, 5}})
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		expect := -1
		actual := slidingPuzzle([][]int{{1, 2, 3}, {5, 4, 0}})
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
}
