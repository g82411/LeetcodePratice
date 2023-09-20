package _658_Minimum_Operations_to_Reduce_X_to_Zero

import "math"

func minOperations(nums []int, x int) int {
	// 窗二解，找到最大的窗make sum(nums[l:r]) = x
	n := len(nums)
	currSum := 0
	for _, num := range nums {
		currSum += num
	}
	l := 0
	minWindowSize := math.MaxInt
	for r := 0; r < n; r++ {
		currSum -= nums[r]
		for currSum < x && l <= r {
			currSum += nums[l]
			l++
		}
		if currSum == x {
			minWindowSize = min(minWindowSize, (n-1-r)+l)
		}
	}
	if minWindowSize == math.MaxInt {
		return -1
	}
	return minWindowSize
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// func minOperations(nums []int, x int) int {
//     // 窗一解，找到最大的窗make sum - x = sum(nums[l:r])
//     n := len(nums)
//     sum := 0
//     for _, num := range nums {
//         sum += num
//     }
//     targetSum := sum - x
//     l := 0
//     currSum := 0
//     maxWindowSize := -1
//     for r := 0; r < n; r++ {
//         currSum += nums[r]
//         for currSum > targetSum && l <= r {
//             currSum -= nums[l]
//             l++
//         }
//         if currSum == targetSum {
//             maxWindowSize = max(maxWindowSize, r - l + 1)
//         }
//     }
//     if maxWindowSize == -1 {
//         return -1
//     }
//     return n - maxWindowSize
// }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
