package _29_Majority_Element_II

import "math"

func majorityElement(nums []int) []int {
	// 投票算法
	// 用2Pass 可以找到眾樹，總之是專業的頻率算法
	// 對於n個元素，出現頻率超過1/3的最多最多最多 兩個
	candidate1 := math.MaxInt
	candidate2 := math.MaxInt
	vote1 := 0
	vote2 := 0
	for _, n := range nums {
		// 投票算法，跟candidate 一樣就投票
		if candidate1 != math.MaxInt && candidate1 == n {
			vote1++
		} else if candidate2 != math.MaxInt && candidate2 == n {
			vote2++
		} else if vote1 == 0 {
			candidate1 = n
			vote1++
		} else if vote2 == 0 {
			candidate2 = n
			vote2++
		} else {
			// 投票算法，跟candidate不一樣就扣票
			vote1--
			vote2--
		}
	}
	// 如果有a, b 兩個元素比較多 那必定最後投票會勝出
	var res []int

	vote1 = 0
	vote2 = 0
	for _, n := range nums {
		if candidate1 != math.MaxInt && candidate1 == n {
			vote1++
		}
		if candidate2 != math.MaxInt && candidate2 == n {
			vote2++
		}
	}
	n := len(nums)
	if vote1 > n/3 {
		res = append(res, candidate1)
	}
	if vote2 > n/3 {
		res = append(res, candidate2)
	}
	return res
}
