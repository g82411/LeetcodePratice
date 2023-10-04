package _458_Max_Dot_Product_of_Two_Subsequences

func maxDotProduct(nums1 []int, nums2 []int) int {
	n := len(nums1)
	m := len(nums2)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			dp[i][j] = nums1[i] * nums2[j]
			if i > 0 {
				dp[i][j] = max(dp[i][j], dp[i-1][j])
			}
			if j > 0 {
				dp[i][j] = max(dp[i][j], dp[i][j-1])
			}
			if i > 0 && j > 0 {
				dp[i][j] = max(dp[i][j], max(dp[i-1][j-1], 0)+nums1[i]*nums2[j])
			}
		}
	}
	return dp[n-1][m-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
