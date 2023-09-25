package _863_Maximum_Length_of_Semi_Decreasing_Subarrays

import "math"

type Stack []int

func (s Stack) Len() int {
	return len(s)
}

func (s Stack) IsEmpty() bool {
	return s.Len() == 0
}

func (s Stack) Top() int {
	return s[s.Len()-1]
}

func (s *Stack) Push(e int) {
	*s = append(*s, e)
}

func (s *Stack) Pop() int {
	prev := *s
	top := prev.Top()
	*s = prev[0 : prev.Len()-1]
	return top
}

// 一開始想說找最大值然後找單調棧
// 然後就GG了, 考慮到一個數列長醬子 [2,3,4,1,11,5,6,7,8,9] 炸
// 當然普通的單調棧肯定不行的
// 最大最小貪心也不對 嗚嗚嗚嗚嗚殺了我我是智障
// 考慮到最貪心的作法其實是N - 1 - 0 + 1
// 我們找到最左邊的值作為最小值, 維護一個遞減的STACK 當然這也是單調棧
// 然後從左為右遞增去找
func maxSubarrayLength(nums []int) int {
	stack := Stack{}
	n := len(nums)
	for i := n - 1; i >= 0; i-- {
		// 單調遞減序列
		if stack.IsEmpty() || nums[i] < nums[stack.Top()] {
			stack.Push(i)
		}
	}
	res := 0
	m := math.MinInt
	for i := 0; i < n; i++ {
		// 如果i 超過棧頂
		for !stack.IsEmpty() && i >= stack.Top() {
			stack.Pop()
		}
		// 如果nums[i]超過之前的最大值
		if nums[i] > m {
			// 更新數列最大值
			m = nums[i]
			for !stack.IsEmpty() && nums[stack.Top()] < m {
				res = max(res, stack.Pop()-i+1)
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
