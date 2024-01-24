package _3017_Count_the_Number_of_Houses_at_a_Certain_Distance_II

// 分析毛刺環，一共有三段
// A - 環 - B
// 分開處理
// AA or BB
// AB
// A環 or B環
// 環內任意兩點

func countOfPairs(n int, x int, y int) []int64 {
	if x > y {
		return countOfPairs(n, y, x)
	}
	ret := make([]int64, n+1)
	// 處理一個特例 x - y == 0
	if x == y {
		for t := 1; t <= n; t++ {
			ret[t] = int64((n - t) * 2)
		}
		return ret[1 : n+1]
	}
	circleSize := y - x + 1
	var calculateDistanceInGlitch func(l int)
	var calculateDistanceBetweenGlitch func(a, b int)
	// p, q 分別為環的左 & 右
	var calculateDistanceBetweenGlitchAndCircle func(l, p, q int)
	var calculateDistanceBetweenCircle func()

	calculateDistanceInGlitch = func(l int) {
		// GG段
		// 起點 1 <= i <= l
		// 終點 i + t <= l
		for t := 1; t <= n; t++ {
			lowerBound := 1
			upperBound := l - t
			ret[t] += max(int64(0), int64(upperBound-lowerBound+1))
		}
	}

	calculateDistanceBetweenGlitch = func(a, b int) {
		// GG'段
		// a段 1 <= i <= a
		// b段 a + 3 <= i + t <= a + 2 + b
		// max(1, a + 3 - t) <= i <= min(a, a + 2 + b - t)
		for t := 1; t <= n; t++ {
			lowerBound := int64(max(1, a+3-t))
			upperBound := int64(min(a, a+2+b-t))
			ret[t] += max(int64(0), upperBound-lowerBound+int64(1))
		}
	}

	calculateDistanceBetweenGlitchAndCircle = func(l, p, q int) {
		// G段與C段
		// 分別為Gp段，即通過捷徑的區段
		// 及Gq段，為原路段
		// 對任意一段長度c
		// i belong 1 <= i <= l
		// i + t belong l + 2 <= i + t <= l + 1 + c
		// 入口點獨立考慮
		for t := 1; t <= n; t++ {
			lower := max(int64(1), int64(l+2-t))
			upper := min(int64(l), int64(l+1+p-t))
			ret[t] += max(int64(0), upper-lower+int64(1))
		}
		for t := 1; t <= n; t++ {
			lower := max(int64(1), int64(l+2-t))
			upper := min(int64(l), int64(l+1+q-t))
			ret[t] += max(int64(0), upper-lower+int64(1))
		}
		// 入口點 i = a
		// i + t = a + 1
		// 1 <= i <= a
		// i = a + 1 - t >= 1
		// a >= t
		for t := 1; t <= n; t++ {
			if l < t {
				continue
			}
			ret[t]++
		}
	}

	calculateDistanceBetweenCircle = func() {
		// CC段
		// t <= cSize - t -> 最短路徑
		// 注意 t = cSize - t要另外考慮
		// 2t = cSize每走一個tick都會繞半圈 會重複計算
		for t := 1; t <= n; t++ {
			if t < circleSize-t {
				ret[t] += int64(circleSize * 2)
				continue
			}
			if t == circleSize/2 {
				ret[t] += int64(circleSize)
			}
		}
	}
	calculateDistanceInGlitch(x - 1)                                                                // AA 段
	calculateDistanceInGlitch(n - y)                                                                // BB 段
	calculateDistanceBetweenGlitch(x-1, n-y)                                                        // AB段
	calculateDistanceBetweenGlitchAndCircle(x-1, (circleSize-1)/2, (circleSize-1)-(circleSize-1)/2) // AC段
	calculateDistanceBetweenGlitchAndCircle(n-y, (circleSize-1)/2, (circleSize-1)-(circleSize-1)/2) // BC段

	// 以上這些的反運算
	for i := range ret {
		ret[i] *= int64(2)
	}

	calculateDistanceBetweenCircle()
	return ret[1 : n+1]
}
