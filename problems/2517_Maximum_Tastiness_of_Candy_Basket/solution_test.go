package _517_Maximum_Tastiness_of_Candy_Basket

import "testing"

func TestMaximumTastiness(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		expect := 8
		actual := maximumTastiness([]int{13, 5, 1, 8, 21, 2}, 3)
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		expect := 2
		actual := maximumTastiness([]int{1, 3, 1}, 2)
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
}
