package _90_The_Maze

type Point struct {
	x int
	y int
}

type Queue []Point

func (q *Queue) Pop() Point {
	result := (*q)[0]
	*q = (*q)[1:]
	return result
}

func (q *Queue) Push(p Point) {
	*q = append((*q), p)
}

func (q *Queue) Len() int {
	return len(*q)
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func hasPath(maze [][]int, start []int, destination []int) bool {
	visited := make(map[Point]bool)
	DIRECTION := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	height := len(maze)
	width := len(maze[0])
	var getTargetPoint func(start Point, vector []int) Point
	q := &Queue{}
	startPoint := Point{start[1], start[0]}
	endPoint := Point{destination[1], destination[0]}
	getTargetPoint = func(start Point, vector []int) Point {
		x := start.x
		y := start.y
		for x >= 0 && y >= 0 && x < width && y < height && maze[y][x] != 1 {
			y += vector[0]
			x += vector[1]
		}
		y -= vector[0]
		x -= vector[1]
		return Point{x, y}
	}
	q.Push(startPoint)
	for !q.IsEmpty() {
		s := q.Pop()
		for _, dir := range DIRECTION {
			next := getTargetPoint(s, dir)
			if visited[next] {
				continue
			}
			if next == endPoint {
				return true
			}
			visited[next] = true
			q.Push(next)
		}
	}
	return false
}
