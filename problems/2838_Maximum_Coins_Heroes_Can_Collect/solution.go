package _838_Maximum_Coins_Heroes_Can_Collect

import (
	"sort"
)

type Stage struct {
	power int
	bonus int
}

func maximumCoins(heroes []int, monsters []int, coins []int) []int64 {
	monsterAmount := len(monsters)
	heroAmount := len(heroes)
	stages := make([]Stage, monsterAmount)
	for i := 0; i < monsterAmount; i++ {
		monster := monsters[i]
		coin := coins[i]
		stages[i] = Stage{monster, coin}
	}
	sort.Slice(stages, func(i, j int) bool {
		return stages[i].power < stages[j].power
	})
	prefixSum := make([]int64, monsterAmount+1)
	for i := 1; i <= monsterAmount; i++ {
		prefixSum[i] = prefixSum[i-1] + int64(stages[i-1].bonus)
	}
	res := make([]int64, heroAmount)
	findMinPower := func(power int) int {
		left, right := 0, len(stages)-1
		insertPos := len(stages) // 默认为最后一个，表示所有的stage的power都不大于给定的power

		for left <= right {
			mid := (left + right) / 2
			if stages[mid].power <= power { // 注意这里的修改，现在是 <=
				left = mid + 1
			} else {
				insertPos = mid
				right = mid - 1
			}
		}

		return insertPos
	}
	for i, power := range heroes {
		a := findMinPower(power)
		res[i] = prefixSum[a]
	}
	return res
}
