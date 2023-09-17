package _489_Find_Critical_and_Pseudo_Critical_Edges_in_Minimum_Spanning_Tree

import "testing"

func TestFindCriticalAndPseudoCriticalEdges(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		expect := [][]int{{0, 1}, {2, 3, 4, 5}}
		actual := findCriticalAndPseudoCriticalEdges(5, [][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 2}, {0, 3, 2}, {0, 4, 3}, {3, 4, 3}, {1, 4, 6}})
		if len(expect) != len(actual) {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})

}
