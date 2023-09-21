package _0_Regular_Expression_Matching

import "testing"

func TestIsMatch(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		expect := false
		actual := isMatch("aa", "a")
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		expect := true
		actual := isMatch("aa", "a*")
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
}
