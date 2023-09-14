package _136_Earliest_Possible_Day_of_Full_Bloom

import (
	"testing"
)

// v3 既然一定球的出deadline, 那我從deadline最常開始做

func TestEarliestFullBloom(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		expect := 9
		actual := earliestFullBloom([]int{1, 4, 3}, []int{2, 3, 1})
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		expect := 9
		actual := earliestFullBloom([]int{1, 2, 3, 2}, []int{2, 1, 2, 1})
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
}
