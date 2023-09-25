package _764_is_Array_a_Preorder_of_Some_Binary_Tree

import "testing"

func TestIsPreorder(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		expect := true
		got := isPreorder([][]int{{0, -1}, {1, 0}, {2, 0}, {3, 2}, {4, 2}})
		if got != expect {
			t.Errorf("got %v expect %v", got, expect)
		}
	})
}
