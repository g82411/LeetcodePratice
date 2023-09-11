package _849_Determine_if_a_Cell_Is_Reachable_at_a_Given_Time

func isReachableAtTime(sx int, sy int, fx int, fy int, t int) bool {
	// 思考題，只要t > 最短距離 那就一定到的了
	dx := abs(sx - fx)
	dy := abs(sy - fy)
	minD := min(dx, dy) + abs(dy-dx)
	if minD == 0 {
		if t == 1 {
			return false
		}
		return true
	}
	return t >= minD
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
