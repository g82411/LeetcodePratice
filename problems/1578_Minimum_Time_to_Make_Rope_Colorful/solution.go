package _578_Minimum_Time_to_Make_Rope_Colorful

func minCost(colors string, neededTime []int) int {
	n := len(colors)
	eliminate := 0
	for i := 0; i < n-1; i++ {
		if colors[i] == colors[i+1] {
			minNeed := min(neededTime[i], neededTime[i+1])
			eliminate += minNeed
			if neededTime[i] > neededTime[i+1] {
				neededTime[i], neededTime[i+1] = neededTime[i+1], neededTime[i]
			}
		}
	}
	return eliminate
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
