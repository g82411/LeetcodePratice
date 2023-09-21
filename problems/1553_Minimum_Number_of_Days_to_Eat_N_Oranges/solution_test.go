package _553_Minimum_Number_of_Days_to_Eat_N_Oranges

import "testing"

func TestMinDays(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		expect := 4
		actual := minDays(10)
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		expect := 3
		actual := minDays(6)
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
}
