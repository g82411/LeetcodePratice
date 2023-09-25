package _063_Number_of_Valid_Subarrays

import "math"

// sub array 的左邊元素是最小元素
// [1, 2, 3, 4, -1]
// [1]
// [1, 2]
// [1, 2, 3]
// [1, 2, 3, 4]
// sum(nextSmaller[i] - i) = count

// go 要自己做ds我真的哭死
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

func validSubarrays(nums []int) int {
	// 觸發強制pop
	nums = append(nums, math.MinInt)
	count := 0
	s := Stack{}
	for i, n := range nums {
		for !s.IsEmpty() && n < nums[s.Top()] {
			count += i - s.Pop()
		}
		s.Push(i)
	}
	return count
}
