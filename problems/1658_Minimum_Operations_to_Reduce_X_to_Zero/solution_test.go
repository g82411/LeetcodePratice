package _658_Minimum_Operations_to_Reduce_X_to_Zero

import "testing"

func TestMinOperations(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		expect := 2
		actual := minOperations([]int{1, 1, 4, 2, 3}, 5)
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
}
