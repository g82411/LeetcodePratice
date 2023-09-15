package _584_Min_Cost_to_Connect_All_Points

import "testing"

func TestMinCostConnectPoints(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		expect := 20
		actual := minCostConnectPoints([][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}})
		if expect != actual {
			t.Errorf("expect: %v, actual: %v", expect, actual)
		}
	})
	t.Run("case 2", func(t *testing.T) {
		expect := 18
		actual := minCostConnectPoints([][]int{{3, 12}, {-2, 5}, {-4, 1}})
		if expect != actual {
			t.Errorf("expect: %v, actual: %v", expect, actual)
		}
	})
	t.Run("case 3", func(t *testing.T) {
		expect := 4
		actual := minCostConnectPoints([][]int{{0, 0}, {1, 1}, {1, 0}, {-1, 1}})
		if expect != actual {
			t.Errorf("expect: %v, actual: %v", expect, actual)
		}
	})
	t.Run("case 4", func(t *testing.T) {
		expect := 4000000
		actual := minCostConnectPoints([][]int{{-1000000, -1000000}, {1000000, 1000000}})
		if expect != actual {
			t.Errorf("expect: %v, actual: %v", expect, actual)
		}
	})
}
