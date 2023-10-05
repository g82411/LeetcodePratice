package _130_Minimum_Cost_Tree_From_Leaf_Values

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

func mctFromLeafValues(arr []int) int {
	n := len(arr)
	nextGreater := make([]int, n)
	prevGreater := make([]int, n)
	for i := range arr {
		nextGreater[i] = math.MaxInt
		prevGreater[i] = math.MaxInt
	}
	s := Stack{}
	for i := range arr {
		for !s.IsEmpty() && arr[s.Top()] <= arr[i] {
			nextGreater[s.Top()] = arr[i]
			s.Pop()
		}
		if !s.IsEmpty() {
			prevGreater[i] = arr[s.Top()]
		}
		s.Push(i)
	}
	result := 0
	for i := range arr {
		x := min(nextGreater[i], prevGreater[i])
		if x == math.MaxInt {
			continue
		}
		result += arr[i] * x
	}

	return result
}

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

// 解一 用DP解
//
//func mctFromLeafValues(arr []int) int {
//	sum := 0
//	n := len(arr)
//	largest := make([][]int, n+1)
//	dp := make([][]int, n+1)
//	for i := range dp {
//		l := make([]int, n+1)
//		d := make([]int, n+1)
//		for j := range l {
//			l[j] = math.MinInt / 10
//			d[j] = math.MaxInt / 10
//		}
//		largest[i] = l
//		dp[i] = d
//	}
//	for i, num := range arr {
//		largest[i][i] = num
//		dp[i][i] = num
//		sum += num
//	}
//	for l := 2; l <= n; l++ {
//		for i := 0; i+l <= n; i++ {
//			j := i + l - 1
//			for k := i; k < j; k++ {
//				dp[i][j] = min(dp[i][j],
//					dp[i][k]+dp[k+1][j]+largest[i][k]*largest[k+1][j],
//				)
//				largest[i][j] = max(largest[i][k], largest[k+1][j])
//			}
//		}
//	}
//	return dp[0][n-1] - sum
//}
//
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
//
//func min(a, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//}
