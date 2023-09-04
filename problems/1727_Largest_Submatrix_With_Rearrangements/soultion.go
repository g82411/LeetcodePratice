package _727_Largest_Submatrix_With_Rearrangements

import (
	"math"
	"sort"
)

// andy評語：妙啊
func largestSubmatrix(mat [][]int) int {
	rows := len(mat)
	cols := len(mat[0])
	maxConsecutiveHeight := make([]int, cols)
	res := math.MinInt
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if mat[row][col] == 0 {
				maxConsecutiveHeight[col] = 0
			} else {
				maxConsecutiveHeight[col] += 1
			}
		}
		// 注意這邊要複製出來 不要打亂原本的排序！！
		temp := make([]int, cols)
		copy(temp, maxConsecutiveHeight)
		// rearrange desc
		sort.Slice(temp, func(i, j int) bool {
			return temp[i] > temp[j]
		})
		minHeight := temp[0]
		for col := 0; col < cols; col++ {
			minHeight = min(minHeight, temp[col])
			res = max(res, temp[col]*(col+1))
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
