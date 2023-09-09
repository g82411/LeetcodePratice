package _499_Max_Value_of_Equation

import "testing"

func TestFindMaxValueOfEquation(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		points := [][]int{{1, 3}, {2, 0}, {5, 10}, {6, -10}}
		k := 1
		got := findMaxValueOfEquation(points, k)
		want := 4
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
	t.Run("case 2", func(t *testing.T) {
		points := [][]int{{0, 0}, {3, 0}, {9, 2}}
		k := 3
		got := findMaxValueOfEquation(points, k)
		want := 3
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
}
