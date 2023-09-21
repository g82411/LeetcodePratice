package _653_Sliding_Subarray_Beauty

func getSubarrayBeauty(nums []int, k int, x int) []int {
	// 這題能注意的只有nums的條件限制 在+-50之間 根本就是誘惑我去掃一遍
	n := len(nums)
	var occurence [101]int
	const offset = 50
	l := 0
	var res []int
	for r := 0; r < n; r++ {
		occurence[nums[r]+offset]++
		if (r - l + 1) == k {
			pos := 0
			e := 0
			for i := 0; i < offset; i++ {
				pos += occurence[i]
				if pos >= x {
					e = i - offset
					break
				}
			}
			res = append(res, e)
			occurence[nums[l]+offset]--
			l++
		}
	}
	return res
}
