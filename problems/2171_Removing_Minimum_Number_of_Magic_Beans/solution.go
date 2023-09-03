package _171_Removing_Minimum_Number_of_Magic_Beans

import (
	"math"
	"sort"
)

func minimumRemoval(beans []int) int64 {
	// 從到大小排序
	// for all k ak < ai, remove step => a1 + a2 + ... + ak sum([a1:ak])
	// for all k ak > ai, remove step => (ak - ai) + ak-1 - ai + ... + ai sum([ak:ai]) - ai * (k - i) = sum([ak:ai]) - ai * i + 1
	sum := 0
	for _, b := range beans {
		sum += b
	}
	// sort beans desc
	sort.Slice(beans, func(i, j int) bool {
		return beans[i] > beans[j]
	})
	minStep := math.MaxInt
	for i := 0; i < len(beans); i++ {
		step := sum - beans[i]*(i+1)
		if minStep > step {
			minStep = step
		}
	}
	return int64(minStep)
}
