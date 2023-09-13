package _35_Candy

import "testing"

func TestCandy(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		expect := 5
		actual := candy([]int{1, 0, 2})
		if expect != actual {
			t.Errorf("expect: %v, actual: %v", expect, actual)
		}
	})
	t.Run("case 2", func(t *testing.T) {
		expect := 4
		actual := candy([]int{1, 2, 2})
		if expect != actual {
			t.Errorf("expect: %v, actual: %v", expect, actual)
		}
	})
}
