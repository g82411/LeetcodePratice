package _849_Determine_if_a_Cell_Is_Reachable_at_a_Given_Time

import "testing"

func TestIsReachableAtTime(t *testing.T) {
	t.Run("case 1 ", func(t *testing.T) {
		want := true
		got := isReachableAtTime(2, 4, 7, 7, 6)
		if want != got {
			t.Errorf("got %v want %v given", got, want)
		}
	})

	t.Run("case 1 ", func(t *testing.T) {
		want := false
		got := isReachableAtTime(3, 1, 7, 3, 3)
		if want != got {
			t.Errorf("got %v want %v given", got, want)
		}
	})
}
