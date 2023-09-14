package _136_Earliest_Possible_Day_of_Full_Bloom

import "sort"

type Flower struct {
	PlantTime int
	GrowTime  int
}

// v3 既然一定球的出deadline, 那我從deadline最常開始做

func earliestFullBloom(plantTime []int, growTime []int) int {
	n := len(plantTime)
	plantPlan := make([]Flower, n)
	for i := range plantTime {
		plantPlan[i] = Flower{
			plantTime[i],
			growTime[i],
		}
	}
	sort.Slice(plantPlan, func(i, j int) bool {
		return plantPlan[i].GrowTime > plantPlan[j].GrowTime
	})
	workDays := 0
	ret := 0
	for _, plan := range plantPlan {
		workDays += plan.PlantTime
		ret = max(ret, workDays+plan.GrowTime)
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// type Flower struct {
//     PlantTime int
//     Deadline int
// }

// v2 優化一下
// func earliestFullBloom(plantTime []int, growTime []int) int {
//     n := len(plantTime)
//     l, r := 1, math.MaxInt / 2
//     plantPlan := make([]Flower, n)
//     for i := range plantTime {
//         plantPlan[i] = Flower{
//             plantTime[i],
//             -growTime[i],
//         }
//     }
//     sort.Slice(plantPlan, func(i, j int) bool {
//         return plantPlan[i].Deadline < plantPlan[j].Deadline
//     })
//     var isPlantAble func (minDeadline int) bool
//     isPlantAble = func (minDeadline int) bool {

//         needDays := 0
//         for _, plan := range plantPlan {
//             needDays += plan.PlantTime
//             if needDays > minDeadline + plan.Deadline {
//                 return false
//             }
//         }
//         return true
//     }
//     for l <= r {
//         m := l + (r - l) / 2
//         if isPlantAble(m) {
//             r = m - 1
//         } else {
//             l = m + 1
//         }
//     }
//     return l
// }

// v1 如果找不到解法, 先確定能不能用binary search 硬幹

// func earliestFullBloom(plantTime []int, growTime []int) int {
//     n := len(plantTime)
//     l, r := 1, math.MaxInt / 2
//     var isPlantAble func (minDeadline int) bool
//     isPlantAble = func (minDeadline int) bool {
//         plantPlan := make([]Flower, n)
//         for i := range plantTime {
//             plantPlan[i] = Flower{
//                 plantTime[i],
//                 minDeadline - growTime[i],
//             }
//         }
//         sort.Slice(plantPlan, func(i, j int) bool {
//             return plantPlan[i].Deadline < plantPlan[j].Deadline
//         })
//         needDays := 0
//         for _, plan := range plantPlan {
//             needDays += plan.PlantTime
//             if needDays > plan.Deadline {
//                 return false
//             }
//         }
//         return true
//     }
//     for l <= r {
//         m := l + (r - l) / 2
//         if isPlantAble(m) {
//             r = m - 1
//         } else {
//             l = m + 1
//         }
//     }
//     return l
// }
