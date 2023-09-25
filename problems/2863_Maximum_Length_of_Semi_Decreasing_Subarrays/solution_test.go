package _863_Maximum_Length_of_Semi_Decreasing_Subarrays

import "testing"

func TestMaxSubarrayLength(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		expect := 8
		actual := maxSubarrayLength([]int{7, 6, 5, 4, 3, 2, 1, 6, 10, 11})
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		expect := 6
		actual := maxSubarrayLength([]int{57, 55, 50, 60, 61, 58, 63, 59, 64, 60, 63})
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
}
