package _000_Minimum_Cost_to_Merge_Stones

import "testing"

func TestMergeStones(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		expect := 20
		actual := mergeStones([]int{3, 2, 4, 1}, 2)
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		expect := -1
		actual := mergeStones([]int{3, 2, 4, 1}, 3)
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
	t.Run("Case 3", func(t *testing.T) {
		expect := 25
		actual := mergeStones([]int{3, 5, 1, 2, 6}, 3)
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
	t.Run("Case 4", func(t *testing.T) {
		expect := 0
		actual := mergeStones([]int{3}, 2)
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
}
