package _063_Number_of_Valid_Subarrays

import "testing"

func TestValidSubarrays(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		expect := 11
		actual := validSubarrays([]int{1, 4, 2, 5, 3})
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		expect := 3
		actual := validSubarrays([]int{3, 2, 1})
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
	t.Run("Case 3", func(t *testing.T) {
		expect := 2
		actual := validSubarrays([]int{2, 2, 2})
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
}
