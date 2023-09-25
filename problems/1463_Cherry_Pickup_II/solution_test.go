package _463_Cherry_Pickup_II

import "testing"

func TestCherryPickup(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		expect := 28
		actual := cherryPickup([][]int{{3, 1, 1}, {2, 5, 1}, {1, 5, 5}, {2, 1, 1}})
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		expect := 24
		actual := cherryPickup([][]int{{1, 0, 0, 0, 0, 0, 1}, {2, 0, 0, 0, 0, 3, 0}, {2, 0, 9, 0, 0, 0, 0}, {0, 3, 0, 5, 4, 0, 0}, {1, 0, 2, 3, 0, 0, 6}})
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
}
