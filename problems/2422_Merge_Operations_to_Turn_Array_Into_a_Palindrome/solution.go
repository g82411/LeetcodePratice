package _422_Merge_Operations_to_Turn_Array_Into_a_Palindrome

func minimumOperations(nums []int) int {
	// 特殊窗
	// 以 4,3,2,1,2,3,1來看
	// 我要從左邊合併 3, 1反過來會從右邊合併3, 1
	// => l sum == right sum
	// if l s == r s => 縮窗
	// ls != rs 縮窗 + 取總，並且此時會發生合併
	n := len(nums)
	count := 0
	if n == 1 {
		return count
	}
	l := 0
	r := n - 1
	lsum := nums[l]
	rsum := nums[r]
	for l < r {
		if lsum == rsum {
			l++
			r--
			lsum = nums[l]
			rsum = nums[r]
			continue
		}
		count++
		if lsum > rsum {
			r--
			rsum += nums[r]
			continue
		}
		l++
		lsum += nums[l]
	}
	return count
}
