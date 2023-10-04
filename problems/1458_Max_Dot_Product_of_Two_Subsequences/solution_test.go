package _458_Max_Dot_Product_of_Two_Subsequences

import "testing"

func TestMaxDotProduct(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		got := maxDotProduct([]int{2, 1, -2, 5}, []int{3, 0, -6})
		want := 18
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		got := maxDotProduct([]int{3, -2}, []int{2, -6, 7})
		want := 21
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
}
