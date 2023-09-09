package _33_Largest_BST_Subtree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type SubTreeSubtreeStatus struct {
	MaxValue int
	MinValue int
	Size     int
}

func largestBSTSubtree(root *TreeNode) int {
	var dfs func(root *TreeNode) SubTreeSubtreeStatus
	dfs = func(root *TreeNode) SubTreeSubtreeStatus {
		if root == nil {
			return SubTreeSubtreeStatus{
				math.MinInt,
				math.MaxInt,
				0,
			}
		}
		leftStatus := dfs(root.Left)
		rightStatus := dfs(root.Right)
		if !(leftStatus.MaxValue < root.Val && root.Val < rightStatus.MinValue) {
			return SubTreeSubtreeStatus{
				math.MaxInt,
				math.MinInt,
				max(leftStatus.Size, rightStatus.Size),
			}
		}
		return SubTreeSubtreeStatus{
			min(root.Val, leftStatus.MinValue),
			max(root.Val, rightStatus.MaxValue),
			leftStatus.Size + rightStatus.Size + 1,
		}
	}
	return dfs(root).Size
}
