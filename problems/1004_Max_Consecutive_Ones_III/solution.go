package _004_Max_Consecutive_Ones_III

func longestOnes(nums []int, k int) int {
	//r := 0
	l := 0
	rk := 0
	n := len(nums)
	ans := 0
	for r := 0; r < n; r++ {
		if nums[r] == 0 {
			rk++
		}
		for rk > k {
			if nums[l] == 0 {
				rk--
			}
			l++
		}
		ans = max(ans, r-l+1)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
