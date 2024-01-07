package _1601_Maximum_Number_of_Achievable_Transfer_Requests

func maximumRequests(n int, requests [][]int) int {
	var check func(state int) int
	m := len(requests)
	check = func(s int) int {
		count := 0
		var net [20]int
		for i := 0; i < m; i++ {
			if ((s >> i) & 1) == 1 {
				net[requests[i][0]]++
				net[requests[i][1]]--
				count++
			}
		}
		for i := 0; i < n; i++ {
			if net[i] > 0 {
				return 0
			}
		}
		return count
	}
	//解1 窮舉
	// for state := 1; state < (1 << m); state++ {
	//     ret = max(ret, check(state))
	// }
	// 解2. Gosper's Hack
	for k := m; k >= 1; k-- {
		state := (1 << k) - 1
		for state < (1 << m) {
			if check(state) > 0 {
				return k
			}
			c := state & -state
			r := state + c
			state = (((r ^ state) >> 2) / c) | r
		}
	}

	return 0
}
