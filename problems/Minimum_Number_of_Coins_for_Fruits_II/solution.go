package _Minimum_Number_of_Coins_for_Fruits_II

type Deque []int

func (q Deque) Len() int {
	return len(q)
}

func (q Deque) IsEmpty() bool {
	return q.Len() == 0
}

func (q Deque) Front() int {
	return q[0]
}

func (q Deque) Back() int {
	return q[q.Len()-1]
}

func (q *Deque) Push(x int) {
	*q = append(*q, x)
}

func (q *Deque) PopFront() int {
	prev := *q
	*q = prev[1:]
	return prev.Front()
}

func (q *Deque) PopBack() int {
	prev := *q
	*q = prev[:q.Len()-1]
	return prev.Back()
}

func minimumCoins(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n+1)
	var deque Deque

	// 複合型考題 第一層 dp
	// dp[i][0] the min cost if we pay prices[i] for i's item
	// dp[i][1] the min cost if we won't pay prices[i] for i's item
	// dp[i][0] = min(dp[i-1][0], dp[i-1][1]) + prices[i]
	// 問題在於dp[i][1]
	// dp[i][1] = min(dp[k][0]) for 2 * k >= i
	// 整體時間複雜度O(n^2)
	// 此時有另外一個考點, 從一個窗中找最大值用deque
	for i := range dp {
		dp[i][0] = math.MaxInt / 2
		dp[i][1] = math.MaxInt / 2
	}
	dp[1][0] = prices[0]
	deque.Push(1)
	for i := 2; i <= n; i++ {
		dp[i][0] = min(dp[i-1][0], dp[i-1][1]) + prices[i-1]
		for !deque.IsEmpty() && deque.Front()*2 < i {
			deque.PopFront()
		}
		dp[i][1] = dp[deque.Front()][0]
		for !deque.IsEmpty() && dp[deque.Back()][0] >= dp[i][0] {
			deque.PopBack()
		}
		deque.Push(i)
	}
	return min(dp[n][0], dp[n][1])
}
