package _12_Burst_Balloons

func padding(n []int) []int {
	n = append([]int{1}, n...)
	n = append(n, 1)
	return n
}

func maxCoins(nums []int) int {
	n := len(nums)
	pn := padding(nums)
	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
	}
	for gap := 1; gap <= n; gap++ {
		for l := 1; l <= (n - gap + 1); l++ {
			r := l + gap - 1
			for k := l; k <= r; k++ {
				dp[l][r] = max(dp[l][r], pn[l-1]*pn[k]*pn[r+1]+dp[l][k-1]+dp[k+1][r])
			}
		}
	}
	return dp[1][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
