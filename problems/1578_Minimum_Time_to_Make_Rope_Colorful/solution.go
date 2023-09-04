package _578_Minimum_Time_to_Make_Rope_Colorful

func minCost(colors string, neededTime []int) int {
	n := len(colors)
	res := 0
	for i := 0; i < n; i++ {
		sumOfTheSame := 0
		maxNeed := 0
		for j := i; j < n; j++ {
			if colors[i] == colors[j] {
				maxNeed = max(maxNeed, neededTime[j])
				sumOfTheSame += neededTime[j]

			}
			// only keep max need time same color
			res += sumOfTheSame - maxNeed
			i = j
		}

	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
