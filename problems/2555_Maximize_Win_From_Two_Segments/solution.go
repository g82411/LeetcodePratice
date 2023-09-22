package _555_Maximize_Win_From_Two_Segments

func maximizeWin(p []int, k int) int {
	// 這題是3 pass貪心的第一題
	// 但是仔細看一下題目也不是不能做 吧？
	maxP := p[len(p)-1]
	minP := p[0]
	n := len(p)
	// 沒辦法分兩段 滾蛋
	if maxP-minP <= k {
		return len(p)
	}
	ml2r := make([]int, n)
	mr2l := make([]int, n)
	l := 0
	mlr := 0
	mrl := 0
	// 從左邊往右邊掃
	for r := 0; r < n; r++ {
		for p[r]-p[l] > k {
			l++
		}
		length := r - l + 1
		mlr = max(mlr, length)
		ml2r[r] = mlr
	}
	r := n - 1
	// 從右邊往左邊掃
	for l := n - 1; l >= 0; l-- {
		for p[r]-p[l] > k {
			r--
		}
		length := r - l + 1
		mrl = max(mrl, length)
		mr2l[l] = mrl
	}
	// 找最大不重複的分割點
	ret := 0
	for g := 0; g+1 < n; g++ {
		ret = max(ret, ml2r[g]+mr2l[g+1])
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
