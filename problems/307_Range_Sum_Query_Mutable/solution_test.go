package _07_Range_Sum_Query_Mutable

import "testing"

func TestNumArray(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		nums := []int{1, 3, 5}
		numArray := Constructor(nums)
		got := numArray.SumRange(0, 2)
		want := 9
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
		numArray.Update(1, 2)
		got = numArray.SumRange(0, 2)
		want = 8
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
}
