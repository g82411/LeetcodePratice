package _653_Sliding_Subarray_Beauty

import "testing"

func TestGetSubarrayBeauty(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		expect := []int{-1, -2, -2}
		actual := getSubarrayBeauty([]int{1, -1, -3, -2, 3}, 3, 2)
		if !equal(expect, actual) {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for idx, val := range a {
		if val != b[idx] {
			return false
		}
	}
	return true
}
