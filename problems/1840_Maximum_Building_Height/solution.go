package _840_Maximum_Building_Height

import "sort"

// 一開始想到單調棧 真的是單你媽
// 對於n來說最高可以蓋多高？
// 如果沒有限制，最高可以蓋到n-1
// 如果有限制呢？最高可以蓋到限制＋1
func maxBuilding(n int, restrictions [][]int) int {
	restrictions = append(restrictions, []int{1, 0})
	sort.Slice(restrictions, func(i, j int) bool {
		return restrictions[i][0] < restrictions[j][0]
	})
	m := len(restrictions)
	pos := make([]int, m)
	limit := make([]int, m)
	for i := range pos {
		pos[i] = restrictions[i][0]
		limit[i] = restrictions[i][1]
	}
	realHeight := make([]int, m)
	realHeight[0] = 0
	// 從左往右掃，i 不應該大於i - j + r[j]
	for i := 1; i < m; i++ {
		realHeight[i] = min(limit[i], realHeight[i-1]+pos[i]-pos[i-1])
	}
	// 從右往左掃，i 不應該大於j - i + r[j]
	for i := m - 2; i >= 0; i-- {
		realHeight[i] = min(realHeight[i], realHeight[i+1]+pos[i+1]-pos[i])
	}
	ret := 0
	for i := 1; i < m; i++ {
		peak := realHeight[i] + (realHeight[i-1]-realHeight[i]-(pos[i-1]-pos[i]))/2
		ret = max(ret, peak)
	}
	// 注意注意！！
	// 最後一棟房子在n 如果n >> m 那要考慮最後一洞房子是不是可以蓋到天際
	ret = max(ret, realHeight[m-1]+n-pos[m-1])
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
