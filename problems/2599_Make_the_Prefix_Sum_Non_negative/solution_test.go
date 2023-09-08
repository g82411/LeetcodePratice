package _599_Make_the_Prefix_Sum_Non_negative

import "testing"

func TestMakePrefSumNonNegative(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		expect := 0
		actual := makePrefSumNonNegative([]int{2, 3, -5, 4})
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		expect := 1
		actual := makePrefSumNonNegative([]int{3, -5, -2, 6})
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		expect := 1
		actual := makePrefSumNonNegative([]int{3, -5, -2, 6})
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
}
