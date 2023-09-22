package _30_Course_Schedule_III

import (
	"container/heap"
	"sort"
)

type Course struct {
	Duration int
	Deadline int
}

type PQ []Course

func (q PQ) Len() int {
	return len(q)
}

func (q PQ) Less(i, j int) bool {
	return q[i].Duration > q[j].Duration
}

func (q PQ) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *PQ) Push(x interface{}) {
	*q = append(*q, x.(Course))
}

func (q *PQ) Pop() interface{} {
	prev := *q
	top := prev[prev.Len()-1]
	*q = prev[0 : prev.Len()-1]
	return top
}

func scheduleCourse(c [][]int) int {
	// Deadline 靠前的先做
	sort.Slice(c, func(i, j int) bool {
		return c[i][1] < c[j][1]
	})
	totalDays := 0
	q := PQ{}
	heap.Init(&q)
	for _, cour := range c {
		dur, dead := cour[0], cour[1]
		totalDays += dur
		heap.Push(&q, Course{
			dur,
			dead,
		})
		if totalDays > dead {
			// 負擔太重的課不修
			mostHeavyCourse := heap.Pop(&q).(Course)
			totalDays -= mostHeavyCourse.Duration

		}
	}
	return q.Len()
}
