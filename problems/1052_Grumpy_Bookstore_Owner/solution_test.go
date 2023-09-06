package _052_Grumpy_Bookstore_Owner

import "testing"

func TestMaxSatisfied(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		actual := maxSatisfied([]int{1, 0, 1, 2, 1, 1, 7, 5}, []int{0, 1, 0, 1, 0, 1, 0, 1}, 3)
		expect := 16
		if actual != expect {
			t.Errorf("actual: %v, expect: %v", actual, expect)
		}
	})
	t.Run("case 2", func(t *testing.T) {
		actual := maxSatisfied([]int{4, 10, 10}, []int{1, 1, 0}, 2)
		expect := 24
		if actual != expect {
			t.Errorf("actual: %v, expect: %v", actual, expect)
		}
	})
}
