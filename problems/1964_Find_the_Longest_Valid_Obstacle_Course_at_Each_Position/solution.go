package _964_Find_the_Longest_Valid_Obstacle_Course_at_Each_Position

import "sort"

func longestObstacleCourseAtEachPosition(nums []int) []int {
	// 找最長遞增子序列
	var res []int

	// 最長遞增子序列
	var arr []int

	n := len(nums)
	// LIS原理
	for i := 0; i < n; i++ {
		// 若LIS = [] 則currLIS = [nums[i]]
		// 若LIS = [1] 且nums[i] = 2
		// 則 curr LIS = [1,2]
		if len(arr) == 0 || nums[i] >= arr[len(arr)-1] {
			arr = append(arr, nums[i])
			res = append(res, len(arr))
		} else {
			// 若LIS = [1,5] curr = 4
			// 則newLIS = [1,4]

			// upper_bound的做法
			insertIndex := sort.Search(len(arr), func(k int) bool {
				return arr[k] > nums[i]
			})
			arr[insertIndex] = nums[i]
			res = append(res, insertIndex+1)
		}
	}
	return res
}
