package _47_Shortest_Path_Visiting_All_Nodes

type Next struct {
	point int
	state int
}

type Queue []Next

func (q *Queue) Enqueue(n Next) {
	*q = append(*q, n)
}

func (q *Queue) Dequeue() Next {
	prev := *q
	last := prev[0]
	*q = prev[1:]
	return last
}

func (q Queue) Len() int {
	return len(q)
}

func (q Queue) IsEmpty() bool {
	return q.Len() == 0
}

func shortestPathLength(graph [][]int) int {
	// 這題是BFS, 但最大的問題是同樣的點可以走 但不可以一直走
	// 所以同樣的點可以走一次
	// state 表示這個bit的點我走過
	// visited[node][state] = true => 不走
	n := len(graph)
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, 1<<n)
	}
	queue := Queue{}
	finalState := 1<<n - 1
	// 從每一個點開始走
	for i := 0; i < n; i++ {
		queue.Enqueue(Next{
			i,
			1 << i,
		})
		visited[i][1<<i] = true
	}
	step := 0
	for !queue.IsEmpty() {
		layerSize := queue.Len()
		for layerSize > 0 {
			layerSize--
			curr := queue.Dequeue()
			point, state := curr.point, curr.state
			for _, np := range graph[point] {
				nextState := state | 1<<np
				if nextState == finalState {
					return step + 1
				}
				// 去過不再去
				if visited[np][state] {
					continue
				}
				visited[np][state] = true
				queue.Enqueue(Next{
					np,
					nextState,
				})
			}
		}
		step++
	}
	return 0
}
