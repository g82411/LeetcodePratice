package _251_Number_of_Flowers_in_Full_Bloom

import "sort"

type Diff struct {
	time   int
	flower int
}

type Person struct {
	index  int
	arrive int
}

func fullBloomFlowers(flowers [][]int, people []int) []int {
	time := make(map[int]int)
	for _, flower := range flowers {
		s, e := flower[0], flower[1]
		time[s]++
		time[e+1]--
	}
	var diff []Diff
	for t, f := range time {
		diff = append(diff, Diff{
			t,
			f,
		})
	}
	var ps []Person
	for i, a := range people {
		ps = append(ps, Person{
			i,
			a,
		})
	}
	sort.Slice(diff, func(i, j int) bool {
		return diff[i].time < diff[j].time
	})
	sort.Slice(ps, func(i, j int) bool {
		return ps[i].arrive < ps[j].arrive
	})

	sum := 0
	j := 0
	ret := make([]int, len(ps))
	for i := 0; i < len(ps); i++ {
		for j < len(diff) && diff[j].time <= ps[i].arrive {
			sum += diff[j].flower
			j++
		}
		ret[ps[i].index] = sum
	}
	return ret
}
