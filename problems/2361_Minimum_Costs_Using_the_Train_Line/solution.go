package _361_Minimum_Costs_Using_the_Train_Line

func minimumCosts(regular []int, express []int, expressCost int) []int64 {
	dp := make([][2]int64, len(regular))
	var ret []int64
	for i := 0; i < len(regular); i++ {
		if i == 0 {
			dp[i][0] = int64(regular[i])
			dp[i][1] = int64(express[i] + expressCost)
			ret = append(ret, min64(dp[i][0], dp[i][1]))
			continue
		}
		dp[i][0] = min64(dp[i-1][0], dp[i-1][1]) + int64(regular[i])
		dp[i][1] = min64(dp[i-1][0]+int64(expressCost+express[i]), dp[i-1][1]+int64(express[i]))
		ret = append(ret, min64(dp[i][0], dp[i][1]))
	}
	return ret
}

func min64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
