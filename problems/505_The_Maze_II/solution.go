package _05_The_Maze_II

import "container/heap"

type Point struct {
	X    int
	Y    int
	Cost int
}

type History struct {
	X int
	Y int
}

type PQ []Point

func (q PQ) Len() int {
	return len(q)
}

func (q PQ) Less(i, j int) bool {
	return q[i].Cost < q[j].Cost
}

func (q PQ) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *PQ) Push(edge interface{}) {
	*q = append(*q, edge.(Point))
}

func (q *PQ) Pop() interface{} {
	n := q.Len()
	prev := *q
	item := prev[n-1]
	*q = prev[0 : n-1]
	return item
}

func shortestDistance(maze [][]int, start []int, destination []int) int {
	startPoint := Point{start[1], start[0], 0}
	visited := make(map[History]bool)
	width := len(maze[0])
	height := len(maze)
	DIRECTION := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	var getTargetPoint func(point Point, direction []int) (Point, History)
	getTargetPoint = func(point Point, direction []int) (Point, History) {
		x := point.X
		y := point.Y
		dx := direction[1]
		dy := direction[0]
		steps := 0
		for (x+dx) >= 0 && (x+dx) < width && (y+dy) >= 0 && (y+dy) < height && maze[y+dy][x+dx] == 0 {
			steps++
			x += dx
			y += dy
		}
		return Point{x, y, point.Cost + steps}, History{x, y}
	}
	var pq PQ
	heap.Init(&pq)
	heap.Push(&pq, startPoint)
	for pq.Len() > 0 {
		s := heap.Pop(&pq).(Point)
		if s.X == destination[1] && s.Y == destination[0] {
			return s.Cost
		}
		if visited[History{s.X, s.Y}] {
			continue
		}
		visited[History{s.X, s.Y}] = true
		for _, dir := range DIRECTION {
			next, his := getTargetPoint(s, dir)
			if visited[his] {
				continue
			}
			heap.Push(&pq, next)
		}
	}
	return -1
}
