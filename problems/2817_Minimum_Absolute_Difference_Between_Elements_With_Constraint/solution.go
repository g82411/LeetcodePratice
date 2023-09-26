package _817_Minimum_Absolute_Difference_Between_Elements_With_Constraint

// 我們從i = x開始陸續把 i - x插入sorted set
// 然後開始找他的upper bound & lower bound
// 這題的重點在于golang中如何產生有序集合
// 此時要記得在golang中 sorted set 可以用rbt進行
import "github.com/emirpasic/gods/trees/redblacktree"

func minAbsoluteDifference(nums []int, x int) int {
	rbt := redblacktree.NewWithIntComparator()
	ans := 1 << 30
	for i := x; i < len(nums); i++ {
		rbt.Put(nums[i-x], nil)
		c, _ := rbt.Ceiling(nums[i])
		f, _ := rbt.Floor(nums[i])
		if c != nil {
			ans = min(ans, c.Key.(int)-nums[i])
		}
		if f != nil {
			ans = min(ans, nums[i]-f.Key.(int))
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
