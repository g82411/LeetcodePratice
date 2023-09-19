package _136_Parallel_Courses

import "testing"

func TestMinimumSemesters(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		expect := 3
		actual := minimumSemesters(3, [][]int{{1, 3}, {2, 3}})
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		expect := -1
		actual := minimumSemesters(3, [][]int{{1, 2}, {2, 3}, {3, 1}})
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
}
