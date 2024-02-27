package _3018_Maximum_Number_of_Removal_Queries_That_Can_Be_Processed_I

// 少數的大區間變小區間問題
// 難點是怎麼想到區間型DP
// 特徵1. 計數問題, 給訂一個list根據某種決策 求最大、最小、Count
// 一開始想得是dp[i][k] => 第i個適用k決策 但是dp應該base on what？ -> 這邊就直接看答案了
// 可能還是要回到複習dp 的種類
// 最後知道是用區間ＩＩ型
// dp[i][j] => maximum number of queries that can be processed when cut down to [i:j]
// let k = dp[i][j]
// if nums[i] >= queries[k] -> dp[i+1][j] = dp[i][j] + 1
// if nums[j] >= queries[k] -> dp[i][j-1] = dp[i][j] + 1
// 後來想想找前後刪一個 或從最前面最後面刪掉一個的 應該可以這樣試試看 因為區間每次都會少

func maximumProcessableQueries(nums []int, q []int) int {
	var dp [1005][1005]int
	n := len(nums)
	maxQ := len(q)
	ret := 0
	for l := n - 1; l >= 1; l-- {
		for i := 0; i+l-1 < n; i++ {
			j := i + l - 1
			if i-1 >= 0 {
				k := dp[i-1][j]
				if k < maxQ && nums[i-1] >= q[k] {
					dp[i][j] = max(dp[i][j], k+1)
				} else {
					dp[i][j] = max(dp[i][j], k)
				}
			}
			if j+1 < n {
				k := dp[i][j+1]
				if k < maxQ && nums[j+1] >= q[k] {
					dp[i][j] = max(dp[i][j], k+1)
				} else {
					dp[i][j] = max(dp[i][j], k)
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		if dp[i][i] < maxQ && nums[i] >= q[dp[i][i]] {
			ret = max(ret, dp[i][i]+1)
		} else {
			ret = max(ret, dp[i][i])
		}
	}
	return ret
}
