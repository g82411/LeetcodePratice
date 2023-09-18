package _29_Majority_Element_II

import "testing"

func TestMajorityElementII(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		expect := []int{3}
		actual := majorityElement([]int{3, 2, 3})
		if expect[0] == actual[0] {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
}
