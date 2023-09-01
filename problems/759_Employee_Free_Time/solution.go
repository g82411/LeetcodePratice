package _59_Employee_Free_Time

import "sort"

type Interval struct {
	Start int
	End   int
}

func employeeFreeTime(schedule [][]*Interval) []*Interval {
	var diff [][2]int
	for _, employee := range schedule {
		for _, interval := range employee {
			diff = append(diff, [2]int{interval.Start, 1})
			diff = append(diff, [2]int{interval.End, -1})
		}
	}
	sort.Slice(diff, func(i, j int) bool {
		if diff[i][0] != diff[j][0] {
			return diff[i][0] < diff[j][0]
		}
		return diff[i][1] > diff[j][1]
	})

	var ret []*Interval
	count := 0
	start := -1
	end := -1
	for _, d := range diff {
		count += d[1]
		if count == 0 && d[1] == -1 {
			start = d[0]
		} else if count == 1 && d[1] == 1 && start != -1 {
			end = d[0]
			ret = append(ret, &Interval{start, end})
		}
	}
	return ret
}
