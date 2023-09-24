package _99_Champagne_Tower

func champagneTower(poured int, query_row int, query_glass int) float64 {
	// 第一層一杯
	prev := []float64{float64(poured)}

	for i := 1; i <= query_row; i++ {
		// 每一層多一杯
		cur := make([]float64, i+1)
		for j := 0; j < i; j++ {
			// 上一層的溢出 = 上一層的量 - 1
			overflow := prev[j] - 1

			if overflow > 0 {
				cur[j] += overflow * 0.5
				cur[j+1] += overflow * 0.5
			}
		}
		prev = cur
	}
	return min(1, prev[query_glass])
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
