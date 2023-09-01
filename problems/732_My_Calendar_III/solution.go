package _732_My_Calendar_III

import "sort"

// 發現一個智障的事情, C++ 的 map 是有序的, Golang 的 map 是無序的 所以要額外排過
// C++ 插一個值是O(1) 我是用 slice 排序, 所以是 O(nlogn) Q_Q

type MyCalendarThree struct {
	occupiedTime map[int]int
	time         []int
}

func Constructor() MyCalendarThree {
	return MyCalendarThree{
		occupiedTime: make(map[int]int),
		time:         make([]int, 0),
	}
}

func (this *MyCalendarThree) Book(startTime int, endTime int) int {
	_, startExist := this.occupiedTime[startTime]
	_, endExist := this.occupiedTime[endTime]
	if !startExist {
		this.time = append(this.time, startTime)
	}
	if !endExist {
		this.time = append(this.time, endTime)
	}
	this.occupiedTime[startTime]++
	this.occupiedTime[endTime]--
	sort.Ints(this.time)
	count := 0
	ret := 0
	for _, t := range this.time {
		v, _ := this.occupiedTime[t]
		count += v
		if count > ret {
			ret = count
		}
	}
	return ret
}
