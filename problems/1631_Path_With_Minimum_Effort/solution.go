package _631_Path_With_Minimum_Effort

// 來個bfs 吧

type Queue struct {
	queue []int
}

func (q *Queue) enqueue(x int) {
	q.queue = append(q.queue, x)
}
func (q *Queue) dequeue() int {
	x := q.queue[0]
	q.queue = q.queue[1:]
	return x
}
func (q *Queue) isEmpty() bool {
	return len(q.queue) == 0
}

func minimumEffortPath(heights [][]int) int {
	//v2 binary search & bfs
	row, col := len(heights), len(heights[0])
	left, right := 0, 1000000
	if row == 1 && col == 1 {
		return 0
	}
	for left < right {
		mid := left + (right-left)/2
		if bfs(heights, mid, row, col) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func bfs(heights [][]int, mid, row, col int) bool {
	visited := make([]bool, row*col)
	visited[0] = true
	q := Queue{}
	q.enqueue(0)
	for !q.isEmpty() {
		x := q.dequeue()
		for _, d := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			nx, ny := x/col+d[0], x%col+d[1]
			if nx < 0 || nx >= row || ny < 0 || ny >= col || visited[nx*col+ny] || abs(heights[x/col][x%col]-heights[nx][ny]) > mid {
				continue
			}
			if nx == row-1 && ny == col-1 {
				return true
			}
			visited[nx*col+ny] = true
			q.enqueue(nx*col + ny)
		}
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// union find 解 很有趣但不快
// type UnionFind struct {
// 	uf   []int
// 	rank []int
// }

// func newUnionFind(n int) *UnionFind {
// 	uf := make([]int, n)
// 	rank := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		uf[i] = i
// 	}
// 	return &UnionFind{uf, rank}
// }

// func (uf *UnionFind) find(x int) int {
// 	if uf.uf[x] != x {
// 		uf.uf[x] = uf.find(uf.uf[x])
// 	}
// 	return uf.uf[x]
// }

// func (uf *UnionFind) union(x, y int) {
// 	rx, ry := uf.find(x), uf.find(y)
// 	if rx == ry {
// 		return
// 	}
// 	if uf.rank[rx] < uf.rank[ry] {
// 		uf.uf[rx] = ry
// 	} else if uf.rank[rx] > uf.rank[ry] {
// 		uf.uf[ry] = rx
// 	} else {
// 		uf.uf[rx] = ry
// 		uf.rank[ry]++
// 	}
// }

// func minimumEffortPath(heights [][]int) int {
// 	row, col := len(heights), len(heights[0])
// 	distanceList := make([][3]int, row)
// 	for y := 0; y < row; y++ {
// 		for x := 0; x < col; x++ {
// 			if y > 0 {
// 				distanceList = append(distanceList, [3]int{y*col + x, (y-1)*col + x, abs(heights[y][x] - heights[y-1][x])})
// 			}
// 			if x > 0 {
// 				distanceList = append(distanceList, [3]int{y*col + x, y*col + x - 1, abs(heights[y][x] - heights[y][x-1])})
// 			}
// 		}
// 	}
// 	// sort distanceList by distance asc
// 	sort.Slice(distanceList, func(i, j int) bool {
// 		return distanceList[i][2] < distanceList[j][2]
// 	})
// 	uf := newUnionFind(row * col)
// 	for _, distance := range distanceList {
// 		uf.union(distance[0], distance[1])
// 		if uf.find(0) == uf.find(row*col-1) {
// 			return distance[2]
// 		}
// 	}
// 	return -1
// }

// func abs(x int) int {
// 	if x < 0 {
// 		return -x
// 	}
// 	return x
// }
