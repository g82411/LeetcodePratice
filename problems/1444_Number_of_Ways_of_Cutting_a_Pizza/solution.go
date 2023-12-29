package _1444_Number_of_Ways_of_Cutting_a_Pizza

func ways(pizza []string, K int) int {
	// 欸嘿
	// 分割完之後 上下刀 上面丟掉 左右刀 右邊丟掉
	// J樣子的話 子問題就是 上左丟掉之後 剩下有多少辦法喔唷
	const Mod = int(1e9 + 7)

	rows := len(pizza)
	cols := len(pizza[0])
	var presum [55][55]int
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			// ref: LeetCode 304
			presum[r+1][c+1] = presum[r+1][c] + presum[r][c+1] - presum[r][c]
			if pizza[r][c] == '.' {
				continue
			}
			presum[r+1][c+1]++
		}
	}
	hasAnyAppleInArea := func(l, t, r, b int) int {
		sum := presum[b+1][r+1] - presum[b+1][l] - presum[t][r+1] + presum[t][l]
		if sum > 0 {
			return 1
		}
		return 0
	}

	var cache [55][55][15]int
	for i := 0; i <= rows; i++ {
		for j := 0; j <= cols; j++ {
			for k := 0; k <= K; k++ {
				cache[i][j][k] = -1
			}
		}
	}

	var dp func(r, c, k int) int
	dp = func(r, c, k int) int {
		if k == 0 {
			return hasAnyAppleInArea(c, r, cols-1, rows-1)
		}
		ans := &cache[r][c][k]
		if *ans >= 0 {
			return *ans
		}
		*ans = 0
		for y := r; y < rows-1; y++ {
			*ans = (*ans + hasAnyAppleInArea(c, r, cols-1, y)*dp(y+1, c, k-1)) % Mod
		}

		for x := c; x < cols-1; x++ {
			*ans = (*ans + hasAnyAppleInArea(c, r, x, rows-1)*dp(r, x+1, k-1)) % Mod
		}
		return *ans
	}
	return dp(0, 0, K-1)
}
