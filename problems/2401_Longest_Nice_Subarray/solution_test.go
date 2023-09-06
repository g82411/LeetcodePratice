package _401_Longest_Nice_Subarray

import "testing"

func TestLongestNiceSubarray(t *testing.T) {
	t.Run("case ", func(t *testing.T) {
		actual := longestNiceSubarray([]int{1, 3, 8, 48, 10})
		expect := 3
		if actual != expect {
			t.Errorf("actual: %v, expect: %v", actual, expect)
		}
	})

	t.Run("case ", func(t *testing.T) {
		actual := longestNiceSubarray([]int{3, 1, 5, 11, 13})
		expect := 1
		if actual != expect {
			t.Errorf("actual: %v, expect: %v", actual, expect)
		}
	})
}
