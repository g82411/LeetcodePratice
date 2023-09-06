package _106_Maximum_Fruits_Harvested_After_at_Most_K_Steps

func maxTotalFruits(fruits [][]int, startPos int, k int) int {
	ret := 0
	// 4 strategy
	// 1. from startPos move to left in x step, and back then move to right in y steps
	// 2. from startPos move to right in y steps
	// 3. from startPos move to right in x step, and back then move to left in y steps
	// 4. from startPos move to left in y steps
	// for any x, y, 2x + y < k
	n := len(fruits)
	prefixSum := make([]int, n)
	position := make([]int, n)
	// bound 的實作都要小心越位，沒有api的情況下會越位月到哭
	for i := 0; i < n; i++ {
		prefixSum[i] = fruits[i][1]
		position[i] = fruits[i][0]
		if i > 0 {
			prefixSum[i] += prefixSum[i-1]
		}
	}
	// find fruit that right of startPos and most close to startPos => lower bound of startPos in position
	lowerBound := func() int {
		l, r := 0, n
		for l < r {
			m := l + (r-l)/2
			if startPos > position[m] {
				l = m + 1
			} else {
				r = m
			}
		}
		return l
	}
	searchStart := lowerBound()
	x := 0
	for i := searchStart; i < n; i++ {
		for position[x] <= startPos && position[i]-startPos+2*(startPos-position[x]) > k {
			x++
		}
		ans := prefixSum[i]
		if x > 0 {
			ans -= prefixSum[x-1]
		}
		if position[x] <= startPos || position[i]-startPos <= k {
			ret = max(ret, ans)
		}
	}
	// find fruit that left of startPos and most close to startPos => upper bound of startPos in position
	upperBound := func() int {
		l, r := 0, n
		for l < r {
			m := l + (r-l)/2
			if startPos >= position[m] {
				l = m + 1
			} else {
				r = m
			}
		}
		return l
	}
	searchStart = upperBound() - 1
	x = n - 1
	for i := searchStart; i >= 0; i-- {
		for position[x] >= startPos && startPos-position[i]+2*(position[x]-startPos) > k {
			x--
		}
		ans := prefixSum[x]
		if i > 0 {
			ans -= prefixSum[i-1]
		}
		if position[x] >= startPos || startPos-position[i] <= k {
			ret = max(ret, ans)
		}
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
