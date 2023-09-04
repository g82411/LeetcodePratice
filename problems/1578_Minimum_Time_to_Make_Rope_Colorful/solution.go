package _578_Minimum_Time_to_Make_Rope_Colorful

func minCost(colors string, neededTime []int) int {
	n := len(colors)
	res := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if colors[i] == colors[j] {
				res += min(neededTime[i], neededTime[i+1])
			}
			i = j
		}

	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
