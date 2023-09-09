package _845_Count_of_Interesting_Subarrays

import "testing"

func TestCountInterestingSubarrays(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		expect := int64(3)
		actual := countInterestingSubarrays([]int{3, 2, 4}, 2, 1)
		if expect != actual {
			t.Errorf("expect: %v, actual: %v", expect, actual)
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		expect := int64(2)
		actual := countInterestingSubarrays([]int{3, 1, 9, 6}, 3, 0)
		if expect != actual {
			t.Errorf("expect: %v, actual: %v", expect, actual)
		}
	})
}
