package _962_Count_Subarrays_Where_Max_Element_Appears_at_Least_K_Times

func countSubarrays(nums []int, k int) int64 {
	l := 0
	n := len(nums)
	t := -1
	for _, num := range nums {
		t = max(t, num)
	}
	tCount := 0
	var ret int64
	for r := 0; r < n; r++ {
		if nums[r] == t {
			tCount++
		}
		for tCount >= k {
			if nums[l] == t {
				tCount--
			}
			l++
		}
		// 這邊為什麼要 + l
		// 根據上面的邏輯 找到右邊界之後，每個做邊界都會是subarray 所以要加上l
		ret += int64(l)
	}
	return ret
}
