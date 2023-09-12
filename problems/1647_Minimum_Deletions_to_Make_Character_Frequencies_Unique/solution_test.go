package _647_Minimum_Deletions_to_Make_Character_Frequencies_Unique

import (
	"testing"
)

func TestMinDeletions(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		expect := 0
		actual := minDeletions("aab")
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		expect := 2
		actual := minDeletions("aaabbbcc")
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
	t.Run("Case 3", func(t *testing.T) {
		expect := 2
		actual := minDeletions("ceabaacb")
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
}
