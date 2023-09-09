package _847_Smallest_Number_With_Given_Digit_Product

import "testing"

func TestSmallestNumber(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		n := int64(105)
		got := smallestNumber(n)
		want := "357"
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		n := int64(44)
		got := smallestNumber(n)
		want := "-1"
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})

}
