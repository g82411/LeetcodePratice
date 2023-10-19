package _355_Maximum_Number_of_Books_You_Can_Take

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

func maximumBooks(books []int) int64 {
	var s Stack
	n := len(books)
	dp := make([]int64, n+1)
	ret := int64(0)
	for i := 0; i < n; i++ {
		for !s.IsEmpty() && books[s.Top()] > books[i]-(i-s.Top()) {
			s.Pop()
		}
		// 1 2 7 8 9
		// dp[1] + (7 + 9) * 3 / 2
		if !s.IsEmpty() {
			h := int64(i - s.Top())
			o := int64(books[i])
			dp[i] = dp[s.Top()] + (o+o-h+1)*h/2
		} else {
			// 1 2 3 4 5
			// (5 + 5 - 4 + 1) * 4  / 2
			// or
			// 2 3 4
			// (4 + 4 - 3 + 1) * 2 / 2
			d := int64(min(i+1, books[i]))
			o := int64(books[i])
			dp[i] = (o + o - d + 1) * d / 2
		}
		s.Push(i)
		ret = max(ret, dp[i])
	}
	return ret
}

func max(a, b int64) int64 {
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
