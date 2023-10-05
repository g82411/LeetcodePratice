package _71_Encode_String_with_Shortest_Length

import "testing"

func TestEncode(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		got := encode("aaa")
		want := "aaa"
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		got := encode("aaaaa")
		want := "5[a]"
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
	t.Run("Case 3", func(t *testing.T) {
		got := encode("aaaaaaaaaa")
		want := "10[a]"
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
}
