package _56_Paint_House

import "testing"

func TestMinCost(t *testing.T) {
	t.Run("Default case 1", func(t *testing.T) {
		costs := [][]int{{17, 2, 17}, {16, 16, 5}, {14, 3, 19}}
		actual := minCost(costs)
		expected := 10
		if actual != expected {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})

	t.Run("Default case 2", func(t *testing.T) {
		costs := [][]int{{7, 6, 2}}
		actual := minCost(costs)
		expected := 2
		if actual != expected {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})

}
