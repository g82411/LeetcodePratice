package _2976_Minimum_Cost_to_Convert_String_I

import "math"

// Floyd-Warshall的精要就是，陸續把點新增進去，然後看看有沒有更短的路徑可以走，如果有就更新，沒有就不更新，最後就會得到最短路徑
func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
	var d [26][26]int64
	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			if i == j {
				continue
			}
			d[i][j] = math.MaxInt64 / int64(3)
		}
	}
	// 先把初始的路徑加進去
	// 另外注意這題的小陷阱，同一個路徑可能出現多種cost
	for i := range original {
		src := int(original[i] - 'a')
		dst := int(changed[i] - 'a')
		d[src][dst] = min(d[src][dst], int64(cost[i]))
	}
	// 注意順序
	// 這邊是先把k放在最外層，因為k代表的是中間點
	for k := 0; k < 26; k++ {
		// 把中間點加進去，然後開掃
		for i := 0; i < 26; i++ {
			for j := 0; j < 26; j++ {
				d[i][j] = min(d[i][j], d[i][k]+d[k][j])
			}
		}
	}

	ret := int64(0)
	for i := range source {
		src := int(source[i] - 'a')
		dst := int(target[i] - 'a')
		dis := d[src][dst]
		if dis == math.MaxInt64/3 {
			return int64(-1)
		}
		ret += dis
	}
	return ret
}
