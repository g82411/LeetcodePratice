package _846_Minimum_Edge_Weight_Equilibrium_Queries_in_a_Tree

import "testing"

func TestMinOperationsQueries(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		actual := minOperationsQueries(8, [][]int{{1, 2, 6}, {1, 3, 4}, {2, 4, 6}, {2, 5, 3}, {3, 6, 6}, {3, 0, 8}, {7, 0, 2}}, [][]int{{4, 6}, {0, 4}, {6, 5}, {7, 4}})
		expect := []int{1, 2, 2, 3}
		if len(actual) != len(expect) {
			t.Errorf("except: %v, actual: %v", expect, actual)
		}
		for i := range expect {
			if actual[i] != expect[i] {
				t.Errorf("except: %v, actual: %v", expect, actual)
			}
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		actual := minOperationsQueries(7, [][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 1}, {3, 4, 2}, {4, 5, 2}, {5, 6, 2}}, [][]int{{0, 3}, {3, 6}, {2, 6}, {0, 6}})
		expect := []int{0, 0, 1, 3}
		if len(actual) != len(expect) {
			t.Errorf("except: %v, actual: %v", expect, actual)
		}
		for i := range expect {
			if actual[i] != expect[i] {
				t.Errorf("except: %v, actual: %v", expect, actual)
			}
		}
	})
}
