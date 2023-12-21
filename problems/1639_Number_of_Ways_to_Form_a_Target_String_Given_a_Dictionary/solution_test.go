package _639_Number_of_Ways_to_Form_a_Target_String_Given_a_Dictionary

import "testing"

func TestNumWays(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		expect := 6
		actual := numWays([]string{"acca", "bbbb", "caca"}, "aba")
		if expect != actual {
			t.Errorf("expect %d, actual %d", expect, actual)
		}
	})
	t.Run("case 2", func(t *testing.T) {
		expect := 1
		actual := numWays([]string{"abba", "baab"}, "bab")
		if expect != actual {
			t.Errorf("expect %d, actual %d", expect, actual)
		}
	})
	t.Run("case 3", func(t *testing.T) {
		expect := 4
		actual := numWays([]string{"abcd"}, "abcd")
		if expect != actual {
			t.Errorf("expect %d, actual %d", expect, actual)
		}
	})
}
