package _318_Number_of_Distinct_Roll_Sequences

func distinctSequences(n int) int {
	mod := int(1e9) + 7
	dp := make([][7][7]int, n+1)
	if n == 1 {
		return 6
	}
	cache := make([][7]int, 7)
	gcd := func(a, b int) int {
		if cache[a][b] > 0 {
			return cache[a][b]
		}
		x, y := a, b
		for y != 0 {
			x, y = y, x%y
		}
		cache[a][b] = x
		cache[b][a] = x
		return x
	}
	// 骰子形式 x a b
	// for all a != b != x && gcd(a, b) == 1
	count := int64(0)
	for a := 1; a <= 6; a++ {
		for b := 1; b <= 6; b++ {
			if a == b {
				continue
			}
			if gcd(a, b) > 1 {
				continue
			}
			dp[2][a][b] = 1
			count++
		}
	}
	if n == 2 {
		return int(count)
	}
	ret := 0
	for i := 3; i <= n; i++ {
		// 列舉前兩棵骰子的狀態
		for a := 1; a <= 6; a++ {
			for b := 1; b <= 6; b++ {
				if a == b || gcd(a, b) > 1 {
					continue
				}
				// 列舉i的狀態
				for x := 1; x <= 6; x++ {
					if x == b {
						continue
					}
					// 只要前兩個是X A 第三個都可以塞b 所以相加
					dp[i][a][b] += dp[i-1][x][a] % mod
				}
				if i == n {
					ret += dp[i][a][b] % mod
				}
			}
		}
	}
	return ret % mod
}
