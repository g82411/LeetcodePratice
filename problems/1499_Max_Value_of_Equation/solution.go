package _499_Max_Value_of_Equation

type Deque []int

func (q *Deque) Push(e int) {
	*q = append(*q, e)
}

func (q Deque) Len() int {
	return len(q)
}

func (q Deque) IsEmpty() bool {
	return q.Len() == 0
}

func (q *Deque) PopBack() int {
	n := q.Len()
	prev := *q
	last := prev[n-1]
	prev = prev[:n-1]
	*q = prev
	return last
}

func (q *Deque) PopFront() int {
	prev := *q
	front := prev[0]
	prev = prev[1:]
	*q = prev
	return front
}

func (q Deque) Front() int {
	return q[0]
}

func (q Deque) Back() int {
	n := q.Len()
	return q[n-1]
}

func findMaxValueOfEquation(points [][]int, k int) int {
	ret := 0
	q := Deque{}
	for i := 0; i < len(points); i++ {
		for !q.IsEmpty() && abs(points[q.Front()][0]-points[i][0]) > k {
			q.PopFront()
		}
		if !q.IsEmpty() {
			ret = max(ret, points[q.Front()][1]+points[i][1]+abs(points[q.Front()][0]-points[i][0]))
		}
		for !q.IsEmpty() && (-points[q.Back()][0]+points[q.Back()][1]) < (-points[i][0]+points[i][1]) {
			q.PopBack()
		}
		q.Push(i)
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
