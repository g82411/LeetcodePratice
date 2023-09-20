package _422_Merge_Operations_to_Turn_Array_Into_a_Palindrome

import "testing"

func TestMinimumOperations(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		expect := 2
		actual := minimumOperations([]int{4, 3, 2, 1, 2, 3, 1})
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		expect := 0
		actual := minimumOperations([]int{1, 1, 1, 1})
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})

	t.Run("Case 3", func(t *testing.T) {
		expect := 3
		actual := minimumOperations([]int{1, 10, 100, 1000})
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
}
