package _41_Cherry_Pickup

import "testing"

func TestCherryPickup(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		expect := 5
		actual := cherryPickup([][]int{{0, 1, -1}, {1, 0, -1}, {1, 1, 1}})
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		expect := 0
		actual := cherryPickup([][]int{{1, 1, -1}, {1, -1, 1}, {-1, 1, 1}})
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
}
