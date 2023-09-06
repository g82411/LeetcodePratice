package binarySearch

func findTarget(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)/2
		if nums[m] == target {
			return m
		}
		if nums[m] > target {
			r = m - 1
		} else if nums[m] < target {
			l = m + 1
		}
	}
	return -1
}

func findFirstTarget(nums []int, target int) int {
	l, r := 0, len(nums)-1
	result := -1
	for l <= r {
		m := l + (r-l)/2
		if nums[m] == target {
			r = m - 1
			result = m
		}
		if nums[m] > target {
			r = m - 1
		} else if nums[m] < target {
			l = m + 1
		}
	}
	return result
}

func listLowerBound(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		m := l + (r-l)/2
		if target > nums[m] {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

func listUpperBound(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		m := l + (r-l)/2
		if target >= nums[m] {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
