package _18_race_car

import "math"

func racecar(target int) int {
	dp := make([]int, target+1)
	for i := 1; i <= target; i++ {
		dp[i] = math.MaxInt
	}
	bitLength := func(a int) int {
		length := 0
		for a > 0 {
			a >>= 1
			length++
		}
		return length
	}
	for i := 1; i <= target; i++ {
		n := bitLength(i)
		// max move distance = 2 ^ n - 1
		maxDistance := 1<<n - 1
		if i == maxDistance {
			dp[i] = n
			continue
		}
		// consider 2 ways
		// 1 way is increase to k and stop
		for j := 0; j < n-1; j++ {
			dp[i] = min(dp[i], n+j+1+dp[i-1<<(n-1)+1<<j])
		}
		// the other way we can accelerate over t and back to i
		if 1<<n-1-i < i {
			dp[i] = min(dp[i], n+1+dp[1<<n-1-i])
		}
	}
	return dp[target]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
