package _1626_Best_Team_With_No_Conflicts

import "sort"

type Player struct {
	Age   int
	Score int
}

func bestTeamScore(scores []int, ages []int) int {
	n := len(scores)
	var players []Player
	for i := 0; i < n; i++ {
		players = append(players, Player{
			ages[i],
			scores[i],
		})
	}
	sort.Slice(players, func(i, j int) bool {
		if players[i].Age == players[j].Age {
			return players[i].Score < players[j].Score
		}
		return players[i].Age < players[j].Age
	})

	dp := make([]int, n+1)
	ret := 0
	for i := 0; i < n; i++ {
		dp[i] = players[i].Score
		for j := 0; j < i; j++ {
			if players[i].Age == players[j].Age {
				dp[i] = max(dp[i], dp[j]+players[i].Score)
				continue
			}
			if players[j].Age < players[i].Age && players[j].Score <= players[i].Score {
				dp[i] = max(dp[i], dp[j]+players[i].Score)
			}
		}
		ret = max(ret, dp[i])
	}
	return ret
}
