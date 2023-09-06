package _841_Maximum_Sum_of_Almost_Unique_Subarray

import "testing"

func TestMaxSum(t *testing.T) {
	t.Run("Basic test", func(t *testing.T) {
		expect := int64(18)
		actual := maxSum([]int{2, 6, 7, 3, 1, 7}, 3, 4)
		if actual != expect {
			t.Errorf("expect: %v, actual: %v", expect, actual)
		}
	})
	t.Run("Basic test1", func(t *testing.T) {
		expect := int64(23)
		actual := maxSum([]int{5, 9, 9, 2, 4, 5, 4}, 1, 3)
		if actual != expect {
			t.Errorf("expect: %v, actual: %v", expect, actual)
		}
	})
	t.Run("Zero case ", func(t *testing.T) {
		expect := int64(0)
		actual := maxSum([]int{1, 2, 1, 2, 1, 2, 1}, 3, 3)
		if actual != expect {
			t.Errorf("expect: %v, actual: %v", expect, actual)
		}
	})
}
