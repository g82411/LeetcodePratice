package _3552_Grid_Teleportation_Traversal

import (
	"container/list"
	"math"
)

type Pos struct {
	X int
	Y int
}

func minMoves(mat []string) int {
	m, n := len(mat), len(mat[0])
	directions := []Pos{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	teleport := make(map[byte][]Pos)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if ch := mat[i][j]; ch >= 'A' && ch <= 'Z' {
				teleport[ch] = append(teleport[ch], Pos{i, j})
			}
		}
	}

	distance := make([][]int, m)
	for i := range distance {
		distance[i] = make([]int, n)
		for j := range distance[i] {
			distance[i][j] = math.MaxInt
		}
	}

	queue := list.New()
	queue.PushBack(Pos{0, 0})
	distance[0][0] = 0

	for queue.Len() > 0 {
		cur := queue.Remove(queue.Front()).(Pos)

		if cur.X == m-1 && cur.Y == n-1 {
			return distance[cur.X][cur.Y]
		}

		ch := mat[cur.X][cur.Y]
		if ch >= 'A' && ch <= 'Z' {
			for _, p := range teleport[ch] {
				if distance[p.X][p.Y] > distance[cur.X][cur.Y] {
					distance[p.X][p.Y] = distance[cur.X][cur.Y]
					queue.PushFront(p)
				}
			}
			delete(teleport, ch) // 防止重複 teleport
		}

		for _, d := range directions {
			nx, ny := cur.X+d.X, cur.Y+d.Y
			if nx >= 0 && nx < m && ny >= 0 && ny < n && mat[nx][ny] != '#' {
				if distance[nx][ny] > distance[cur.X][cur.Y]+1 {
					distance[nx][ny] = distance[cur.X][cur.Y] + 1
					queue.PushBack(Pos{nx, ny})
				}
			}
		}
	}

	return -1
}
