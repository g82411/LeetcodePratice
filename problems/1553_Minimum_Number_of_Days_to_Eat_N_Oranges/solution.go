package _553_Minimum_Number_of_Days_to_Eat_N_Oranges

func minDays(n int) int {
	// 答案其實很不重要 重要的是時間複雜度
	cache := make(map[int]int)
	var dp func(x int) int
	dp = func(x int) int {
		if x == 0 {
			return 0
		}
		if x == 1 {
			return 1
		}
		if x == 2 || x == 3 {
			return 2
		}
		if _, e := cache[x]; e {
			return cache[x]
		}
		ans := 1 + min(x%2+dp(x/2), x%3+dp(x/3))
		cache[x] = ans
		return ans
	}
	// 問題在於它的時間複雜度是log^n
	//        f(n)
	//      f(n/2)  f(n/3)
	//    f(n/2/2) f(n/2/3)...
	// 所以她其實是一個數型走訪
	return dp(n)

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
