package _862_Maximum_Element_Sum_of_a_Complete_Subset_of_Indices

import "testing"

func TestMaximumSum(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		expect := int64(16)
		actual := maximumSum([]int{8, 7, 3, 5, 7, 2, 4, 9})
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		expect := int64(19)
		actual := maximumSum([]int{5, 10, 3, 10, 1, 13, 7, 9, 4})
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
}
