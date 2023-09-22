package _787_Ways_to_Express_an_Integer_as_Sum_of_Powers

import "testing"

func TestNumberOfWays(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		got := numberOfWays(10, 2)
		want := 1
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		got := numberOfWays(4, 1)
		want := 2
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
}
