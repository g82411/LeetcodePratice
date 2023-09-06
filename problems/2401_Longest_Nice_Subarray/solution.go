package _401_Longest_Nice_Subarray

func longestNiceSubarray(nums []int) int {
	l, r := 0, 0
	mask := 0
	n := len(nums)
	length := 0
	for r < n {
		if nums[r]&mask == 0 {
			mask += nums[r]
			length = max(length, r-l+1)
			r++
			continue
		}
		if l < r {
			mask -= nums[l]
			l++
		}
	}
	return length
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
