package _147_Number_of_Ways_to_Divide_a_Long_Corridor

func numberOfWays(corridor string) int {
	// 乍看？？？？
	// 題目好長
	// 椅子之間是盆栽
	//
	var pos []int
	for i := range corridor {
		if corridor[i] == 'S' {
			pos = append(pos, i)
		}
	}
	totalSeat := len(pos)
	if totalSeat%2 > 0 {
		return 0
	}
	if totalSeat == 0 {
		return 0
	}

	ret := 1
	for i := 2; i+2 <= totalSeat; i += 2 {
		diff := pos[i] - pos[i-1]
		ret *= diff
		ret %= int(1e9 + 7)
	}
	return ret
}
