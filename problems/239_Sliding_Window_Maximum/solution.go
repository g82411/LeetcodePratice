package _39_sliding_window_max

import "math"

type Dequeue []int

func (q Dequeue) Len() int {
	return len(q)
}

func (q Dequeue) IsEmpty() bool {
	return q.Len() == 0
}

func (q *Dequeue) PushBack(n int) {
	*q = append(*q, n)
}

func (q *Dequeue) PopBack() int {
	if q.IsEmpty() {
		return math.MinInt
	}
	prev := *q
	back := prev[q.Len()-1]
	*q = prev[:prev.Len()-1]
	return back
}

func (q *Dequeue) PopFront() int {
	if q.IsEmpty() {
		return -1
	}
	prev := *q
	front := prev[0]
	*q = prev[1:]
	return front
}

func (q Dequeue) Back() int {
	if q.IsEmpty() {
		return -1
	}
	return q[q.Len()-1]
}

func (q Dequeue) Front() int {
	if q.IsEmpty() {
		return -1
	}
	return q[0]
}

// disable copilot
func maxSlidingWindow(nums []int, k int) []int {
	//一個窗口只有一個最大值 -> 找next large
	q := Dequeue{}
	ans := []int{}
	for i := range nums {
		// monotonic queue
		for !q.IsEmpty() && nums[q.Back()] < nums[i] {
			q.PopBack()
		}
		q.PushBack(i)
		// max value over window size
		if q.Front() <= (i - k) {
			q.PopFront()
		}
		if i+1 >= k {
			ans = append(ans, nums[q.Front()])
		}
	}
	return ans
}
