package _554_Strings_Differ_by_One_Character

import "testing"

func TestDifferByOne(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		actual := differByOne([]string{"abcd", "acbd", "aacd"})
		if !actual {
			t.Errorf("Expect true")
		}
	})
}
