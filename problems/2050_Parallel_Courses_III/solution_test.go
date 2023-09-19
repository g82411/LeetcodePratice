package _050_Parallel_Courses_III

import "testing"

func TestMinimumTime(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		expect := 8
		actual := minimumTime(3, [][]int{{1, 3}, {2, 3}}, []int{3, 2, 5})
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		expect := 12
		actual := minimumTime(5, [][]int{{1, 5}, {2, 5}, {3, 5}, {3, 4}, {4, 5}}, []int{1, 2, 3, 4, 5})
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
}
