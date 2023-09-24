package _800_Shortest_String_That_Contains_Three_Strings

import "testing"

/*
 */
func TestMinimumString(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		expect := "aaabca"
		actual := minimumString("abc", "bca", "aaa")
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
}
