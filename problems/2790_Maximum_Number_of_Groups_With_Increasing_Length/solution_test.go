package _790_Maximum_Number_of_Groups_With_Increasing_Length

import "testing"

func TestMaxIncreasingGroups(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		expect := 3
		actual := maxIncreasingGroups([]int{1, 2, 5})
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		expect := 2
		actual := maxIncreasingGroups([]int{2, 1, 2})
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
	t.Run("Case 3", func(t *testing.T) {
		expect := 1
		actual := maxIncreasingGroups([]int{1, 1})
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})

}
