package _747_Count_Zero_Request_Servers

import "testing"

func TestCountServers(t *testing.T) {
	//n = 3, logs = [[1,3],[2,6],[1,5]], x = 5, queries = [10,11]
	//n = 3, logs = [[2,4],[2,1],[1,2],[3,1]], x = 2, queries = [3,4]
	t.Run("Default test case1", func(t *testing.T) {
		actual := countServers(3, [][]int{{1, 3}, {2, 6}, {1, 5}}, 5, []int{10, 11})
		expected := []int{1, 2}
		if len(actual) != len(expected) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})

	t.Run("Default test case1", func(t *testing.T) {
		actual := countServers(3, [][]int{{2, 4}, {2, 1}, {1, 2}, {3, 1}}, 2, []int{3, 4})
		expected := []int{0, 1}
		if len(actual) != len(expected) {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})
}
