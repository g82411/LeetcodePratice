package _868_The_Wording_Game

import "testing"

func TestCanAliceWin(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		expect := false
		actual := canAliceWin([]string{"avokado", "dabar"}, []string{"brazil"})
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		expect := true
		actual := canAliceWin([]string{"ananas", "atlas", "banana"}, []string{"albatros", "cikla", "nogomet"})
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})

	t.Run("Case 3", func(t *testing.T) {
		expect := true
		actual := canAliceWin([]string{"hrvatska", "zastava"}, []string{"bijeli", "galeb"})
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
}
