package _106_Maximum_Fruits_Harvested_After_at_Most_K_Steps

import "testing"

func TestMaxTotalFruits(t *testing.T) {
	t.Run("Test max strategy is move single way", func(t *testing.T) {
		expect := 9
		fruits := [][]int{{2, 8}, {6, 3}, {8, 6}}
		startPos := 5
		k := 4
		actual := maxTotalFruits(fruits, startPos, k)
		if actual != expect {
			t.Errorf("except: %v, actual: %v", expect, actual)
		}
	})

	t.Run("Test need move bidirectional to get max", func(t *testing.T) {
		expect := 14
		fruits := [][]int{{0, 9}, {4, 1}, {5, 7}, {6, 2}, {7, 4}, {10, 9}}
		startPos := 5
		k := 4
		actual := maxTotalFruits(fruits, startPos, k)
		if actual != expect {
			t.Errorf("except: %v, actual: %v", expect, actual)
		}
	})

	t.Run("Test max value = 0 ", func(t *testing.T) {
		expect := 0
		fruits := [][]int{{0, 3}, {6, 4}, {8, 5}}
		startPos := 3
		k := 2
		actual := maxTotalFruits(fruits, startPos, k)
		if actual != expect {
			t.Errorf("except: %v, actual: %v", expect, actual)
		}
	})

	t.Run("Test max value = 0 ", func(t *testing.T) {
		expect := 0
		fruits := [][]int{{0, 10000}}
		startPos := 200000
		k := 0
		actual := maxTotalFruits(fruits, startPos, k)
		if actual != expect {
			t.Errorf("except: %v, actual: %v", expect, actual)
		}
	})

	t.Run("Test max value = 0 ", func(t *testing.T) {
		expect := 10000
		fruits := [][]int{{200000, 10000}}
		startPos := 0
		k := 200000
		actual := maxTotalFruits(fruits, startPos, k)
		if actual != expect {
			t.Errorf("except: %v, actual: %v", expect, actual)
		}
	})
}
