package _514_Path_with_Maximum_Probability

import "testing"

func TestMaxProbability(t *testing.T) {
	t.Run("Basic testcase", func(t *testing.T) {
		expect := 0.25000
		actual := maxProbability(3, [][]int{{0, 1}, {1, 2}, {0, 2}}, []float64{0.5, 0.5, 0.2}, 0, 2)
		if actual != expect {
			t.Errorf("except: %v, actual: %v", expect, actual)
		}
	})

	t.Run("Basic testcase", func(t *testing.T) {
		expect := 0.30000
		actual := maxProbability(3, [][]int{{0, 1}, {1, 2}, {0, 2}}, []float64{0.5, 0.5, 0.3}, 0, 2)
		if actual != expect {
			t.Errorf("except: %v, actual: %v", expect, actual)
		}
	})
}
