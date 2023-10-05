package _478_Allocate_Mailboxes

import "testing"

func TestMinDistance(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		got := minDistance([]int{1, 4, 8, 10, 20}, 3)
		want := 5
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		got := minDistance([]int{2, 3, 5, 12, 18}, 2)
		want := 9
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
}
