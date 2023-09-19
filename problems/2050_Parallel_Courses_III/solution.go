package _050_Parallel_Courses_III

type Queue []int

func (q Queue) Len() int {
	return len(q)
}
func (q Queue) IsEmpty() bool {
	return q.Len() == 0
}

func (q *Queue) Enqueue(e int) {
	*q = append(*q, e)
}

func (q *Queue) Dequeue() int {
	prev := *q
	front := prev[0]
	*q = prev[1:]
	return front
}

func minimumTime(n int, relations [][]int, time []int) int {
	// kahn
	inDegree := make([]int, n+1)
	next := make(map[int][]int)
	maxTime := make([]int, n+1)
	for _, rel := range relations {
		src, dst := rel[0], rel[1]
		inDegree[dst]++
		next[src] = append(next[src], dst)
	}
	q := Queue{}
	for i := 1; i <= n; i++ {
		if inDegree[i] == 0 {
			q.Enqueue(i)
			maxTime[i] = time[i-1]
		}
	}
	// 這邊有億點點陷阱
	// 如果上一層有一個點14 一個點4 4下面有一個點3
	// 則按照一般kahn算法 總時間會變成14 + 3 = 17
	// 但是實際上修完4 就可以修3 所以總時間應該是14
	// ret := 0
	// for !q.IsEmpty() {
	//     maxTime := 0
	//     nextQueue := Queue{}
	//     for _, node := range q {
	//         maxTime = max(maxTime, time[node-1])
	//         for _, nextCourse := range next[node] {
	//             inDegree[nextCourse] --
	//             if inDegree[nextCourse] == 0 {
	//                 nextQueue.Enqueue(nextCourse)
	//             }
	//         }
	//     }
	//     ret += maxTime
	//     q = nextQueue
	// }
	for !q.IsEmpty() {
		node := q.Dequeue()
		for _, neighbor := range next[node] {
			maxTime[neighbor] = max(maxTime[neighbor], maxTime[node]+time[neighbor-1])
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				q.Enqueue(neighbor)
			}
		}
	}
	ret := 0
	for _, mt := range maxTime {
		ret = max(ret, mt)
	}
	return ret

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
