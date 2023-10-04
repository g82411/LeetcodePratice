package _46_Remove_Boxes

func removeBoxes(boxes []int) int {
	n := len(boxes)
	var memo [105][105][105]int
	var dfs func(l, r, k int) int
	dfs = func(l, r, k int) int {
		if l > r {
			return 0
		}
		if memo[l][r][k] != 0 {
			return memo[l][r][k]
		}
		for r > l && boxes[r] == boxes[r-1] {
			r--
			k++
		}
		if memo[l][r][k] != 0 {
			return memo[l][r][k]
		}
		memo[l][r][k] = (k+1)*(k+1) + dfs(l, r-1, 0)
		for i := l; i < r; i++ {
			if boxes[i] == boxes[r] {
				memo[l][r][k] = max(memo[l][r][k], dfs(l, i, k+1)+dfs(i+1, r-1, 0))
			}
		}
		return memo[l][r][k]
	}
	return dfs(0, n-1, 0)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
