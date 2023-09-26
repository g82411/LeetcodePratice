package _547_Minimum_Cost_to_Cut_a_Stick

import "testing"

func TestMinCost(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		expect := 9
		actual := minCost(7, []int{1, 3, 4, 5})
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		expect := 11
		actual := minCost(9, []int{5, 6, 1, 4, 2})
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
}
