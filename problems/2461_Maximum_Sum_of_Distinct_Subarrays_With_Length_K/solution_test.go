package _461_Maximum_Sum_of_Distinct_Subarrays_With_Length_K

import "testing"

func TestMaximumSubarraySum(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		actual := maximumSubarraySum([]int{1, 5, 4, 2, 9, 9, 9}, 3)
		expect := int64(15)
		if actual != expect {
			t.Errorf("except: %v, actual: %v", expect, actual)
		}
	})

	t.Run("Case 1", func(t *testing.T) {
		actual := maximumSubarraySum([]int{4, 4, 4}, 3)
		expect := int64(0)
		if actual != expect {
			t.Errorf("except: %v, actual: %v", expect, actual)
		}
	})
}
