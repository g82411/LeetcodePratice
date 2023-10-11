package _036_Maximum_Alternating_Subarray_Sum

import "testing"

func TestMaximumAlternatingSubarraySum(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		got := maximumAlternatingSubarraySum([]int{4, 2, 5, 3})
		want := int64(7)
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
}
