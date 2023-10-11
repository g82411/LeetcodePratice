package _039_Minimum_Score_Triangulation_of_Polygon

import "testing"

func TestMinScoreTriangulation(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		got := minScoreTriangulation([]int{1, 2, 3})
		want := 6
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
	t.Run("case 2", func(t *testing.T) {
		got := minScoreTriangulation([]int{3, 7, 4, 5})
		want := 144
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
}
