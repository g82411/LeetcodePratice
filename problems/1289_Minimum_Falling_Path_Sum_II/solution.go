package _289_Minimum_Falling_Path_Sum_II

func minFallingPathSum(m [][]int) int {
	minInLastRow := 101
	lastInLastRow := 101
	minIndex := -1
	height := len(m)
	width := len(m[0])
	for i := 0; i < width; i++ {
		if minInLastRow == 101 {
			minInLastRow = m[0][i]
			minIndex = i
			continue
		}
		if minInLastRow > m[0][i] {
			lastInLastRow = minInLastRow
			minInLastRow, minIndex = m[0][i], i
		}
	}
	for i := 1; i < height; i++ {
		minInThisRow := 101
		lastInThisRow := 101
		minThisIndex := -1
		for j := 0; j < width; j++ {
			if minInThisRow == 101 {
				if j == minIndex {
					minInThisRow = m[i][j] + lastInLastRow
				} else {
					minInLastRow = m[i][j] + minInLastRow
				}
				minThisIndex = j
				continue
			}
			sum := m[i][j] + minInLastRow
			if j == minIndex {
				sum = m[i][j] + lastInLastRow
			}
			if sum < minInThisRow {
				lastInThisRow = minInThisRow
				minInThisRow, minThisIndex = sum, j
			}
			minInLastRow, minIndex, lastInLastRow = minInThisRow, minThisIndex, lastInThisRow
		}
	}
	return minInLastRow
}
