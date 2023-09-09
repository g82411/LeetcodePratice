package _33_Largest_BST_Subtree

import "testing"

func TestLargestBSTSubtree(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		root := &TreeNode{Val: 10}
		root.Left = &TreeNode{Val: 5}
		root.Right = &TreeNode{Val: 15}
		root.Left.Left = &TreeNode{Val: 1}
		root.Left.Right = &TreeNode{Val: 8}
		root.Right.Right = &TreeNode{Val: 7}
		got := largestBSTSubtree(root)
		want := 3
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
	t.Run("case 2", func(t *testing.T) {
		root := &TreeNode{Val: 4}
		root.Left = &TreeNode{Val: 2}
		root.Right = &TreeNode{Val: 7}
		root.Left.Left = &TreeNode{Val: 2}
		root.Left.Right = &TreeNode{Val: 3}
		root.Right.Left = &TreeNode{Val: 5}
		root.Right.Right = &TreeNode{Val: 8}
		root.Right.Left.Left = &TreeNode{Val: 4}
		root.Right.Left.Right = &TreeNode{Val: 6}
		got := largestBSTSubtree(root)
		want := 2
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
}
