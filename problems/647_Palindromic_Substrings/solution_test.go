package _47_Palindromic_Substrings

import "testing"

func TestCountSubstrings(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		expect := 3
		actual := countSubstrings("abc")
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		expect := 6
		actual := countSubstrings("aaa")
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})

}
