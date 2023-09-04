package _56_Paint_House

func minCost(costs [][]int) int {
	startWithR := costs[0][0]
	startWithG := costs[0][1]
	startWithB := costs[0][2]
	for i, cost := range costs {
		if i == 0 {
			continue
		}
		costR, costG, costB := cost[0], cost[1], cost[2]
		tR, tG, tB := startWithR, startWithG, startWithB
		startWithR = min(tG, tB) + costR
		startWithG = min(tR, tB) + costG
		startWithB = min(tR, tG) + costB
	}
	return min(startWithR, min(startWithG, startWithB))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
