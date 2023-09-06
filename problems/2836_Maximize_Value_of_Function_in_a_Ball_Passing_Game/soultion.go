package _836_Maximize_Value_of_Function_in_a_Ball_Passing_Game

import (
	"math"
)

func getMaxFunctionValue(receiver []int, k int64) int64 {
	n := len(receiver)
	ancestors := make([][35]int, 10005)
	dp := make([][35]int64, 10005)
	maxLifting := int(math.Ceil(math.Log2(float64(k))))
	var buildAncestorTable func()
	var binaryLiftingK func() []int
	buildAncestorTable = func() {
		for i := 0; i < n; i++ {
			ancestors[i][0] = receiver[i]
			dp[i][0] = int64(receiver[i])
		}
		for j := 1; j <= maxLifting; j++ {
			for i := 0; i < n; i++ {
				ancestors[i][j] = ancestors[ancestors[i][j-1]][j-1]
				dp[i][j] = dp[i][j-1] + dp[ancestors[i][j-1]][j-1]
			}
		}
	}
	binaryLiftingK = func() []int {
		bits := []int{}
		for i := 0; i <= maxLifting; i++ {
			if k>>i&1 == 1 {
				bits = append(bits, i)
			}
		}
		return bits
	}
	buildAncestorTable()
	bitMap := binaryLiftingK()
	ret := int64(0)
	for i := 0; i < n; i++ {
		ans := int64(i)
		cur := i
		for _, bit := range bitMap {
			ans += dp[cur][bit]
			cur = ancestors[cur][bit]
		}
		ret = max(ret, ans)
	}
	return ret
}
func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
