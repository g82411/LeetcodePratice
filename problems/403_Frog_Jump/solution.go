package _03_Frog_Jump

func canCross(stones []int) bool {
	n := len(stones)
	memo := make(map[[2]int]bool)
	stoneSet := make(map[int]bool)
	isStonePos := func(pos int) bool {
		if _, e := stoneSet[pos]; e {
			return true
		}
		return false
	}
	for _, stone := range stones {
		stoneSet[stone] = true
	}

	var dfs func(curr, lastJump int) bool

	dfs = func(curr, lastJump int) bool {
		if curr == stones[n-1] {
			return true
		}
		if !isStonePos(curr) {
			return false
		}
		if _, e := memo[[2]int{curr, lastJump}]; e {
			return false
		}
		if lastJump > 1 && dfs(curr+lastJump-1, lastJump-1) {
			return true
		}
		if lastJump > 0 && dfs(curr+lastJump, lastJump) {
			return true
		}
		if dfs(curr+lastJump+1, lastJump+1) {
			return true
		}
		memo[[2]int{curr, lastJump}] = false
		return false
	}

	return dfs(0, 0)
}
