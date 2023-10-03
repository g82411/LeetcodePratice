package _334_Subarray_With_Elements_Greater_Than_Varying_Threshold

import "math"

type Stack []int

func (s Stack) Len() int {
	return len(s)
}

func (s Stack) Top() int {
	return s[s.Len()-1]
}

func (s Stack) IsEmpty() bool {
	return s.Len() == 0
}

func (s *Stack) Push(e int) {
	*s = append(*s, e)
}

func (s *Stack) Pop() int {
	prev := *s
	top := prev.Top()
	prev = prev[0 : prev.Len()-1]
	*s = prev
	return top
}

// func largestRectangleArea(heights []int) int {
//     // one pass, 對於stack內 next smaller 是i
//     // 那prev smaller 明顯會是 stack的第二個元素
//     // 插一個最小的元素0 避免獸人永不pop
//     heights = append(heights, 0)
//     heights = append([]int{0}, heights...)

//     // heights = append(heights, 0)
//     stack := &Stack{}
//     ans := 0
//     for i := range heights {
//         // 問題來了 如果這邊height本來就是遞增的 那就獸人永不pop了
//         for !stack.IsEmpty() && heights[stack.Top()] > heights[i] {
//             h := heights[stack.Pop()]
//             prevSamller := stack.Top()
//             nextSmaller := i
//             area := (nextSmaller - prevSamller - 1) * h
//             ans = max(ans, area)
//         }
//         stack.Push(i)
//     }

//     return ans
// }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func validSubarraySize(heights []int, threshold int) int {
	heights = append(heights, 0)
	heights = append([]int{0}, heights...)

	// heights = append(heights, 0)
	stack := &Stack{}
	ans := math.MaxInt
	for i := range heights {
		// 問題來了 如果這邊height本來就是遞增的 那就獸人永不pop了
		for !stack.IsEmpty() && heights[stack.Top()] > heights[i] {
			h := heights[stack.Pop()]
			prevSamller := stack.Top()
			nextSmaller := i
			area := (nextSmaller - prevSamller - 1) * h
			if area > threshold {
				ans = min(ans, nextSmaller-prevSamller-1)
			}
		}
		stack.Push(i)
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans
}
