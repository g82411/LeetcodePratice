package _866_Beautiful_Towers_II

// 最佳解是什麼？
// max{山的體積}
// 山的體積 = e + 上坡的體積 + 下坡的體積
// 對於山坡i考慮兩種狀況
/*
              *
              *
        *     *
   *    *     *
   *    *     *
  i-2 i - 1 i
*/
// 若 i 為最高點，上坡體積(i)我們只需要求0~i-1 + i
/*
        *
   *    *
   *    *
   *    *     *
   *    *     *
  i-2 i - 1 i
*/
// 反之 若山坡i 為最低點 我們就需要整平
// 將i 到prevSmaller(i)之間變成i
// 設prevSmaller j area = area[j] + (j - i + 1) * height[i]
// 對於上坡則是nextSmaller
// 然後這題跟上一題2865一饃一樣 差別在於測資範圍

type Stack []int

func (s Stack) Len() int {
	return len(s)
}

func (s Stack) IsEmpty() bool {
	return s.Len() == 0
}

func (s *Stack) Push(e int) {
	*s = append(*s, e)
}

func (s Stack) Top() int {
	return s[s.Len()-1]
}

func (s *Stack) Pop() int {
	prev := *s
	top := prev.Top()
	*s = prev[0 : prev.Len()-1]
	return top
}

func maximumSumOfHeights(maxHeights []int) int64 {
	n := len(maxHeights)
	stack := Stack{}
	prevSmaller := make([]int, n)
	nextSmaller := make([]int, n)
	for i := range maxHeights {
		prevSmaller[i] = -1
		nextSmaller[i] = n

	}
	for i, h := range maxHeights {
		for !stack.IsEmpty() && maxHeights[stack.Top()] > h {
			idx := stack.Pop()
			nextSmaller[idx] = i
		}
		// 考慮stack 剩下[1,2,3] 則h[3] 就是h[i]的prevSmaller
		if !stack.IsEmpty() {
			prevSmaller[i] = stack.Top()
		}
		stack.Push(i)
	}
	uphillArea := make([]int64, n)
	downhillArea := make([]int64, n+1)
	for i := 0; i < n; i++ {
		prevSmall := prevSmaller[i]
		if prevSmall == -1 {
			// i是最小值 整平
			uphillArea[i] = int64((i + 1) * maxHeights[i])
		} else {
			uphillArea[i] = int64((i-prevSmall)*maxHeights[i]) + uphillArea[prevSmall]
		}
	}
	for i := n - 1; i >= 0; i-- {
		nextSmall := nextSmaller[i]
		if nextSmall == n {
			// 下坡最小值
			downhillArea[i] = int64((n - i) * maxHeights[i])
		} else {
			downhillArea[i] = downhillArea[nextSmall] + int64((nextSmall-i)*maxHeights[i])
		}
	}
	ans := int64(0)
	for i := 0; i < n; i++ {
		ans = max64(ans, uphillArea[i]+downhillArea[i+1])
	}
	return ans
}

func max64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
