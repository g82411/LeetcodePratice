package _861_Maximum_Number_of_Alloys

// 二分搜尋法之題目越繞越多彎

// 用第i個機器create alloys 需要composition[i]個各式金屬
// 最好最好的狀況 一個合金 只需要一種金屬 並且這種金屬只要一塊錢
// 所以你有多少budge + stock 就可以買多少金屬

func maxNumberOfAlloys(n int, k int, budget int, composition [][]int, stock []int, cost []int) int {
	l := 0
	r := stock[0] + budget
	for _, s := range stock {
		r = min(r, s+budget)
	}
	canCreateAlloys := func(machine, numsOfAlloys int) bool {
		needMoney := 0
		com := composition[machine]
		// 能夠生產n塊金屬 com[i] * n > s or 買完要不超過預算
		for i := 0; i < n; i++ {
			inventure := stock[i]
			requirement := com[i] * numsOfAlloys
			if inventure >= requirement {
				continue
			}
			short := requirement - inventure
			needMoney += short * cost[i]
			if needMoney > budget {
				return false
			}
		}
		return true
	}
	ans := 0
	for mac := range composition {
		ml := l
		mr := r
		for ml <= mr {
			mid := ml + (mr-ml)/2
			if canCreateAlloys(mac, mid) {
				ml = mid + 1
			} else {
				mr = mid - 1
			}
		}
		ans = max(ans, ml)
	}
	return ans - 1

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
