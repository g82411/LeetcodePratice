package _39_sliding_window_max

import "testing"

func TestMaxSlidingWindow(t *testing.T) {
	// test case
	t.Run("Case 1", func(t *testing.T) {
		nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
		k := 3
		got := maxSlidingWindow(nums, k)
		want := []int{3, 3, 5, 5, 6, 7}
		if len(got) != len(want) {
			t.Errorf("got %v want %v given", got, want)
		}
		for i := range want {
			if got[i] != want[i] {
				t.Errorf("got %v want %v given", got, want)
			}
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		nums := []int{1}
		k := 1
		got := maxSlidingWindow(nums, k)
		want := []int{1}
		if len(got) != len(want) {
			t.Errorf("got %v want %v given", got, want)
		}
		for i := range want {
			if got[i] != want[i] {
				t.Errorf("got %v want %v given", got, want)
			}
		}
	})

}
