package _99_The_Maze_III

import "container/heap"

type Direction struct {
	DX   int
	DY   int
	Path string
}

type Point struct {
	X    int
	Y    int
	Path string
	Cost int
}

type PQ []Point

func (q PQ) Len() int {
	return len(q)
}

func (q PQ) IsEmpty() bool {
	return len(q) == 0
}

func (q PQ) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q PQ) Less(i, j int) bool {
	if q[i].Cost == q[j].Cost {
		return q[i].Path < q[j].Path
	}
	return q[i].Cost < q[j].Cost
}

func (q *PQ) Push(edge interface{}) {
	*q = append(*q, edge.(Point))
}

func (q *PQ) Pop() interface{} {
	prev := *q
	top := prev[q.Len()-1]
	*q = prev[0 : q.Len()-1]
	return top
}

func findShortestWay(maze [][]int, ball []int, hole []int) string {
	visited := make(map[[2]int]bool)
	width := len(maze[0])
	height := len(maze)
	DIRECTION := []Direction{
		Direction{1, 0, "r"},
		Direction{-1, 0, "l"},
		Direction{0, 1, "d"},
		Direction{0, -1, "u"},
	}
	var getTargetPoint func(s Point, d Direction) Point
	start := Point{
		ball[1],
		ball[0],
		"",
		0,
	}
	maze[hole[0]][hole[1]] = 9
	var pq PQ
	heap.Init(&pq)
	heap.Push(&pq, start)
	getTargetPoint = func(s Point, d Direction) Point {
		x := s.X
		y := s.Y
		dx := d.DX
		dy := d.DY
		dir := d.Path
		steps := 0

		for (x+dx) >= 0 && (x+dx) < width && (y+dy) >= 0 && (y+dy) < height && (maze[y+dy][x+dx] == 0 || maze[y+dy][x+dx] == 9) {
			steps++
			x += dx
			y += dy
			if maze[y][x] == 9 {
				break
			}
		}
		return Point{
			x,
			y,
			s.Path + dir,
			s.Cost + steps,
		}
	}
	for !pq.IsEmpty() {
		top := heap.Pop(&pq).(Point)
		if visited[[2]int{top.X, top.Y}] {
			continue
		}
		if top.X == hole[1] && top.Y == hole[0] {
			return top.Path
		}
		visited[[2]int{top.X, top.Y}] = true
		for _, d := range DIRECTION {
			next := getTargetPoint(top, d)
			nextCoordinate := [2]int{next.X, next.Y}
			if visited[nextCoordinate] {
				continue
			}
			heap.Push(&pq, next)
		}
	}
	return "impossible"

}
