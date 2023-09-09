package _845_Count_of_Interesting_Subarrays

// TODO: 多練練
// 一步一步來 首先要找到一個nums[i] % m == k
// 然後這一步的cnt % module = k
// cnt的意思就是到這一步的長度 所以如果到這一步的長度是k1 則所有k1 - k 的組合都可以變成ret

func countInterestingSubarrays(nums []int, modulo int, k int) int64 {
	count := 0
	n := len(nums)
	cntMap := make(map[int]int64)
	ret := int64(0)
	cntMap[0]++
	for i := 0; i < n; i++ {
		if nums[i]%modulo == k {
			count++
		}
		k1 := count % modulo
		remainder := (k1 - k + modulo) % modulo
		ret += cntMap[remainder]
		cntMap[k1]++
	}
	return ret
}
