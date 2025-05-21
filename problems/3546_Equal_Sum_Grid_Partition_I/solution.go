package _3546_Equal_Sum_Grid_Partition_I

func canPartitionGrid(grid [][]int) bool {
	m, n := len(grid), len(grid[0])
	rowSum := make([]int64, m)
	colSum := make([]int64, n)
	total := int64(0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			rowSum[i] += int64(grid[i][j])
			colSum[j] += int64(grid[i][j])
			total += int64(grid[i][j])
		}
	}
	if total%2 != 0 {
		return false
	}
	target := total / 2
	var prev int64
	for i := 0; i < m; i++ {
		prev += rowSum[i]
		if prev == target {
			return true
		}
	}
	prev = int64(0)
	for j := 0; j < n; j++ {
		prev += colSum[j]
		if prev == target {
			return true
		}
	}
	return false
}
