package _770_Maximum_Score_from_Performing_Multiplication_Operations

import "testing"

func TestMaximumScore(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		expect := 14
		actual := maximumScore([]int{1, 2, 3}, []int{3, 2, 1})
		if expect != actual {
			t.Errorf("expect %v actual %v", expect, actual)
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		expect := 102
		actual := maximumScore([]int{-5, -3, -3, -2, 7, 1}, []int{-10, -5, 3, 4, 6})
		if expect != actual {
			t.Errorf("expect %v actual %v", expect, actual)
		}
	})

}
