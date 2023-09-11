package _223_Sum_of_Scores_of_Built_Strings

import "testing"

func TestSumScores(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		expect := int64(9)
		actual := sumScores("babab")
		if expect != actual {
			t.Errorf("Expect: %v, actual: %v", expect, actual)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		expect := int64(14)
		actual := sumScores("azbazbzaz")
		if expect != actual {
			t.Errorf("Expect: %v, actual: %v", expect, actual)
		}
	})
}
