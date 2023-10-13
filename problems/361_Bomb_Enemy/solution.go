package _61_Bomb_Enemy

// 同一列 如果沒卡牆 那be 都會一樣
// 如果卡牆 那就從牆後開始算

func maxKilledEnemies(grid [][]byte) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])
	rowCount := 0
	res := 0
	colCount := make([]int, cols)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if c == 0 || grid[r][c-1] == 'W' {
				rowCount = 0
				for k := c; k < cols && grid[r][k] != 'W'; k++ {
					if grid[r][k] == 'E' {
						rowCount++
					}
				}
			}
			if r == 0 || grid[r-1][c] == 'W' {
				colCount[c] = 0
				for k := r; k < rows && grid[k][c] != 'W'; k++ {
					if grid[k][c] == 'E' {
						colCount[c]++
					}
				}
			}
			if grid[r][c] == '0' {
				res = max(res, colCount[c]+rowCount)
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
