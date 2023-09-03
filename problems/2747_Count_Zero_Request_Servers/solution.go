package _747_Count_Zero_Request_Servers

import "sort"

func countServers(n int, logs [][]int, x int, queries []int) []int {
	// sorted query & logs by time
	sort.Slice(logs, func(i, j int) bool {
		return logs[i][1] < logs[j][1]
	})
	var queryWithIndex [][2]int
	for idx, query := range queries {
		queryWithIndex = append(queryWithIndex, [2]int{query, idx})
	}
	sort.Slice(queryWithIndex, func(i, j int) bool {
		return queryWithIndex[i][0] < queryWithIndex[j][0]
	})
	l, r := 0, 0
	loadingServers := make(map[int]int)
	ret := make([]int, len(queries))
	for _, query := range queryWithIndex {
		time, idx := query[0], query[1]
		// fill loading server
		for r < len(logs) && logs[r][1] <= time {
			serverId := logs[r][0]
			loadingServers[serverId]++
			r++
		}
		// flush no loading server
		for l < len(logs) && logs[l][1] < time-x {
			serverId := logs[l][0]
			loadingServers[serverId]--
			// if serverId has no more loading, flush out
			if loadingServers[serverId] <= 0 {
				delete(loadingServers, serverId)
			}
			l++
		}
		ret[idx] = n - len(loadingServers)
	}
	return ret
}
