package _517_Maximum_Tastiness_of_Candy_Basket

import (
	"math"
	"sort"
)

func maximumTastiness(price []int, k int) int {
	// 這邊二分搜尋還是看得出來的
	// 問題在於如何構造check並使其單調

	// 首先這是可排序的!!!!!!!! -> 單調
	// 故找一個最小的間隔d 找k個不小於d的間隔

	var check func(diff int) bool
	l := 0
	r := math.MaxInt / 2
	sort.Ints(price)

	check = func(diff int) bool {
		count := 1
		for i := 0; i < len(price); i++ {
			j := i
			for j < len(price) && price[j]-price[i] < diff {
				j++
			}
			if j < len(price) {
				count++
				i = j - 1
			} else {
				break
			}
		}
		return count >= k
	}
	for l < r {
		m := r - (r-l)/2
		if check(m) {
			l = m
			continue
		}
		r = m - 1
	}
	return l

}
