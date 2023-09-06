package _25_Contiguous_Array

func findMaxLength(nums []int) int {
	firstPosition := make(map[int]int)
	sum := 0
	ans := 0
	for i, num := range nums {
		addition := 1
		if num == 0 {
			addition = -1
		}
		sum += addition
		if sum == 0 {
			ans = max(ans, i+1)
			continue
		}
		position, exist := firstPosition[sum]
		if !exist {
			firstPosition[sum] = i
			continue
		}
		ans = max(ans, i-position)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
