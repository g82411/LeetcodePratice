package _267_Check_if_There_Is_a_Valid_Parentheses_String_Path

// 迷宮反正就是from dp[i-1][j], dp[j-1][i] to dp[i][j]沒問題
// 但問題是dp表達怎麼寫
// 考慮到總共有200個左右括號 可以寫成var dp [i][j][k]bool
// 表示從0,0 ~ i,j 是否存在k個未分配的左括號
// so if grid[i][j] = '(' then dp[i][j][k] = dp[i][j-1][k-1] || dp[i-1][j][k-1]
// else dp[i][j][k] = dp[i][j-1][k+1] || dp[i-1][j][k+1]
func hasValidPath(grid [][]byte) bool {
	var dp [101][101][103]bool

	rows := len(grid)
	cols := len(grid[0])
	if grid[0][0] == '(' {
		dp[0][0][1] = true
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			for k := 0; k <= (rows+cols-1)/2; k++ {
				if i == 0 && j == 0 {
					continue
				}
				if k > 0 && grid[i][j] == '(' {
					dp[i][j][k] = (j > 0 && dp[i][j-1][k-1]) || (i > 0 && dp[i-1][j][k-1])
				} else if grid[i][j] == ')' {
					dp[i][j][k] = (j > 0 && dp[i][j-1][k+1]) || (i > 0 && dp[i-1][j][k+1])

				}
			}
		}
	}
	return dp[rows-1][cols-1][0]
}
