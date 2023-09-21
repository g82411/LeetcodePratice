package _793_Maximum_Score_of_a_Good_Subarray

import "testing"

func TestMaximumScore(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		expect := 4
		actual := maximumScore([]int{1, 4, 3, 7, 4, 5}, 3)
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		expect := 6
		actual := maximumScore([]int{5, 5, 4, 5, 4, 1, 1, 1}, 0)
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
}
