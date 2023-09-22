package _30_Course_Schedule_III

import "testing"

func TestScheduleCourse(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		got := scheduleCourse([][]int{{100, 200}, {200, 1300}, {1000, 1250}, {2000, 3200}})
		want := 3
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		got := scheduleCourse([][]int{{1, 2}})
		want := 1
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
}
