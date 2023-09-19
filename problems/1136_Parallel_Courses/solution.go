package _136_Parallel_Courses

type Queue []int

func (q Queue) Len() int {
	return len(q)
}

func (q Queue) IsEmpty() bool {
	return q.Len() == 0
}

func (q *Queue) Push(e int) {
	*q = append(*q, e)
}

func (q *Queue) Pop() int {
	prev := *q
	front := prev[0]
	*q = prev[1:]
	return front
}

func minimumSemesters(n int, relations [][]int) int {
	// 能修到所有的課 = 不存在環
	// kahn 演算法
	// 找到所有node的入度
	inDegree := make([]int, n+1)
	e := make(map[int][]int)
	for _, rel := range relations {
		src, dst := rel[0], rel[1]
		inDegree[dst]++
		e[src] = append(e[src], dst)
	}
	// 將入度=0的node作為起點
	q := Queue{}
	for i := 1; i <= n; i++ {
		if inDegree[i] == 0 {
			q.Push(i)
		}
	}
	step := 0
	ableStudyCourses := 0
	for !q.IsEmpty() {
		step++
		newQueue := Queue{}
		for _, node := range q {
			ableStudyCourses++
			for _, next := range e[node] {
				// 將下一個點的入度 --
				inDegree[next]--
				// 如果入度=0 代表這是最後一個可以進入next的點
				if inDegree[next] == 0 {
					newQueue.Push(next)
				}
			}
		}
		q = newQueue
	}
	if ableStudyCourses == n {
		return step
	}
	return -1
}
