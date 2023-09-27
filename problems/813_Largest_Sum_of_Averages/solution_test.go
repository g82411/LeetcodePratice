package _13_Largest_Sum_of_Averages

import "testing"

func TestLargestSumOfAverages(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		expect := 20.0
		actual := largestSumOfAverages([]int{9, 1, 2, 3, 9}, 3)
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		expect := 24.0
		actual := largestSumOfAverages([]int{1, 2, 3, 4, 5, 6, 7}, 4)
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
}
