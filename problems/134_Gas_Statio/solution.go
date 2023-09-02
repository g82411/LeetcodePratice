package _34_Gas_Statio

func canCompleteCircuit(gas []int, cost []int) int {
	// 難度??
	// 從不虧的那一格開始走
	// 走完沒油就return -1
	globalRemainGas := 0
	remainGas := 0
	start := 0
	n := len(gas)
	for i := 0; i < n; i++ {
		g := gas[i]
		c := cost[i]
		globalRemainGas = globalRemainGas + g - c
		remainGas = remainGas + g - c
		if remainGas < 0 {
			remainGas = 0
			start = i + 1
		}

	}
	if globalRemainGas < 0 {
		return -1
	}
	return start
}
