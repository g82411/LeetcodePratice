package _964_Find_the_Longest_Valid_Obstacle_Course_at_Each_Position

import "testing"

func TestLongestObstacleCourseAtEachPosition(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		got := longestObstacleCourseAtEachPosition([]int{1, 2, 3, 2})
		want := []int{1, 2, 3, 3}
		if !equal(got, want) {
			t.Errorf("got %v want %v given", got, want)
		}
	})
}

func equal([]int, []int) bool {
	return true
}
