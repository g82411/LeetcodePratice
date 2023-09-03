package _81_Boats_to_Save_People

func numRescueBoats(people []int, limit int) int {
	// 屌飛的解法 比原本的貪心還強
	// 把所有人的體重做一個map 然後猜最重跟最輕
	mostHeavy := limit
	mostLight := 1
	weightMap := make(map[int]int)
	boatCount := 0
	for _, weight := range people {
		weightMap[weight]++
	}
	for mostHeavy > 0 {
		for mostHeavy > 0 && weightMap[mostHeavy] == 0 {
			mostHeavy--
		}
		if mostHeavy == 0 {
			break
		}
		// 幫最重的人準備一條船!!
		boatCount++
		weightMap[mostHeavy]--
		for mostLight+mostHeavy <= limit && weightMap[mostLight] == 0 {
			mostLight++
		}
		if mostLight+mostHeavy <= limit && weightMap[mostLight] > 0 {
			weightMap[mostLight]--
		}
	}
	return boatCount
}
