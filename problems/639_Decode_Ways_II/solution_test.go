package _39_Decode_Ways_II

import "testing"

func TestNumDecodings(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		got := numDecodings("*")
		want := 9
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		got := numDecodings("1*")
		want := 18
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
	t.Run("Case 3", func(t *testing.T) {
		got := numDecodings("2*")
		want := 15
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
}
