package _578_Minimum_Time_to_Make_Rope_Colorful

import "testing"

func TestMinCost(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		colors := "abc"
		neededTime := []int{1, 2, 3}
		expect := 0
		actual := minCost(colors, neededTime)
		if actual != expect {
			t.Errorf("colors: %v, neededTime: %v, expect: %v, actual: %v", colors, neededTime, expect, actual)
		}
	})
	t.Run("test2", func(t *testing.T) {
		colors := "aabaa"
		neededTime := []int{1, 2, 3, 4, 1}
		expect := 2
		actual := minCost(colors, neededTime)
		if actual != expect {
			t.Errorf("colors: %v, neededTime: %v, expect: %v, actual: %v", colors, neededTime, expect, actual)
		}
	})
}
