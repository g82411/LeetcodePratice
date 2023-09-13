package _35_Candy

func candy(ratings []int) int {
	// 很妙，從左往右看一次，再從右往左看一次
	n := len(ratings)
	d := make([]int, n)
	// 每個人先發一顆
	for i := range d {
		d[i]++
	}
	// 看自己左邊的人
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			d[i] = max(1, d[i-1]+1)
		}
	}
	// 看右邊
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			d[i] = max(d[i], d[i+1]+1)
		}
	}
	sum := 0
	for i := range d {
		sum += d[i]
	}
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
