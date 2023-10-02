package _876_Count_Visited_Nodes_in_a_Directed_Graph

type Queue []int

func (q Queue) Len() int {
	return len(q)
}

func (q Queue) IsEmpty() bool {
	return q.Len() == 0
}

func (q Queue) Front() int {
	return q[0]
}

func (q *Queue) Pop() int {
	prev := *q
	front := prev[0]
	*q = prev[1:]
	return front
}

func (q *Queue) Push(e int) {
	*q = append(*q, e)
}

func countVisitedNodes(edges []int) []int {
	n := len(edges)
	inDegree := make([]int, n)
	for _, edge := range edges {
		inDegree[edge]++
	}
	q := Queue{}
	for i := 0; i < n; i++ {
		if inDegree[i] == 0 {
			q.Push(i)
		}
	}

	for !q.IsEmpty() {
		size := q.Len()
		for size > 0 {
			curr := q.Pop()
			neighbor := edges[curr]
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				q.Push(neighbor)
			}
			size--
		}
	}

	// Count cycle length as length

	res := make([]int, n)
	for i := 0; i < n; i++ {
		if inDegree[i] == 0 {
			continue
		}
		if res[i] != 0 {
			continue
		}

		// count cycle length
		cycleLength := 1
		pointer := i
		for edges[pointer] != i {
			cycleLength++
			pointer = edges[pointer]
		}
		// fill legnth in cycle
		pointer = i
		for edges[pointer] != i {
			res[pointer] = cycleLength
			pointer = edges[pointer]
		}
	}

	var dfs func(curr int) int
	dfs = func(curr int) int {
		if res[curr] != 0 {
			return res[curr]
		}
		res[curr] = dfs(edges[curr]) + 1
		return res[curr]
	}
	for i := 0; i < n; i++ {
		if inDegree[i] != 0 {
			continue
		}
		dfs(i)
	}

	return res
}
