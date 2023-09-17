package _856_Minimum_Array_Length_After_Pair_Removals

// 做貪心的題目要學會怎麼建模跟建立testcase
// 設出現最多的元素是k 次數是m
// 則 1 2 3 4 4 5 k k k k k k k k k k k k k k 6 7 8 9
// 若 k 可以抵消比 k 大與比k小的數字
// 則若k <= len(nums) / 2
// 則利用剩下的數字可以互相抵銷
// consider: 1 2 3 4 5 6 K K K K K K K K 8 8 8 8 8 9 9 9 10 10
// 可以用10 & 9抵銷 9剩1個 K 可以抵銷剩下的 8 & 9
// 然後
// consider: 1 2 3 4 5  K K K K K K K K 8 8 8 8 8 9 9 9 10 10
// 考慮到這會剩下最後1個
// 最後若 k > len(nums) / 2
// 全部都抵銷會剩下k 所以數字會剩下 m - (n - m)個 => 2 * m - n

func minLengthAfterRemovals(nums []int) int {
	n := len(nums)
	frequency := make(map[int]int)
	for _, num := range nums {
		frequency[num]++
	}
	maxFrequency := 0
	for _, freq := range frequency {
		maxFrequency = max(maxFrequency, freq)
	}
	if maxFrequency <= n/2 {
		if n%2 == 1 {
			return 1
		}
		return 0
	}
	ret := 2*maxFrequency - n
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
