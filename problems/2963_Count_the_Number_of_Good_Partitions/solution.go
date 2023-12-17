package _963_Count_the_Number_of_Good_Partitions

func numberOfGoodPartitions(nums []int) int {
	// 插分數列姐姐姐
	n := len(nums)
	left := make(map[int]int)
	right := make(map[int]int)
	for i := 0; i < n; i++ {
		right[nums[i]] = i
	}
	for i := n - 1; i >= 0; i-- {
		left[nums[i]] = i
	}
	diff := make([]int, n)
	for num, _ := range left {
		diff[left[num]]++
		diff[right[num]]--
	}

	sum := 0
	count := 0
	for i := 0; i < n; i++ {
		sum += diff[i]
		if sum == 0 {
			count++
		}
	}
	ret := 1

	for i := 0; i < count-1; i++ {
		ret = ret * 2 % int(1e9+7)
	}
	return ret
}
