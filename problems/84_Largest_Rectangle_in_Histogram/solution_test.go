package _4_Largest_Rectangle_in_Histogram

import "testing"

func TestLargestRectangleArea(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		expect := 10
		actual := largestRectangleArea([]int{2, 1, 5, 6, 2, 3})
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		expect := 4
		actual := largestRectangleArea([]int{2, 4})
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
}
