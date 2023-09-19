package _494_Parallel_Courses_II

import "testing"

func TestMinNumberOfSemesters(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		expect := 3
		actual := minNumberOfSemesters(4, [][]int{{2, 1}, {3, 1}, {1, 4}}, 2)
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		expect := 3
		actual := minNumberOfSemesters(5, [][]int{{2, 1}, {3, 1}, {4, 1}, {1, 5}}, 2)
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
}
