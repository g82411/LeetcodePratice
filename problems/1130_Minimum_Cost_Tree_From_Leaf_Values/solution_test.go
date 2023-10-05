package _130_Minimum_Cost_Tree_From_Leaf_Values

import "testing"

func TestMctFromLeafValues(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		got := mctFromLeafValues([]int{6, 2, 4})
		want := 32
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		got := mctFromLeafValues([]int{7, 12, 8, 10})
		want := 284
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
}
