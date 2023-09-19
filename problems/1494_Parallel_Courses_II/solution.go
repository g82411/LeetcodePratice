package _494_Parallel_Courses_II

import (
	"math"
	"math/bits"
)

func minNumberOfSemesters(n int, relations [][]int, k int) int {
	// 狀態壓縮dp做 每一步都是bit mask
	dp := make([]int, 1<<n)
	for i := range dp {
		dp[i] = math.MaxInt / 2
	}
	preState := make([]int, 1<<n)
	preReq := make([]int, n)
	// 這邊先算出所有的prevCourse
	for _, rel := range relations {
		prev, next := rel[0], rel[1]
		preReq[next-1] |= 1 << (prev - 1)
	}
	// 這邊要算State, state 要計算所有先修集合 e.g state = 1000雖然preCoure是0001 但是他前面還有0010 & 0100
	// 所以state 4 = 1111
	for state := 0; state < (1 << n); state++ {
		for i := 0; i < n; i++ {
			if (state>>i)&1 == 1 {
				preState[state] |= preReq[i]
			}
		}
		// 沒有先修課 && 修課沒有超過學分上限
		if preState[state] == 0 && bits.OnesCount32(uint32(state)) <= k {
			dp[state] = 1
		}
	}
	dp[0] = 0

	for state := 1; state < 1<<n; state++ {
		for subState := state; subState > 0; subState = (subState - 1) & state {

			// 三個條件
			// 1. 修課不超過學分上限
			// 2. 先修課必須要完修
			if bits.OnesCount32(uint32(state))-bits.OnesCount32(uint32(subState)) > k {
				continue
			}
			// preState = 1111100
			// subState = 1111110
			// &=         1111100
			if preState[state]&subState != preState[state] {
				continue
			}
			dp[state] = min(dp[state], dp[subState]+1)

		}

	}
	return dp[1<<n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
