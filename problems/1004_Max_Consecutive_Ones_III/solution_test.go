package _004_Max_Consecutive_Ones_III

import "testing"

func TestLongestOnes(t *testing.T) {
	t.Run("simple test", func(t *testing.T) {
		nums := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
		k := 2
		expect := 6
		actual := longestOnes(nums, k)
		if actual != expect {
			t.Errorf("expect: %v, actual: %v", expect, actual)
		}
	})

	t.Run("simple test", func(t *testing.T) {
		nums := []int{0, 0, 0, 0, 0}
		k := 2
		expect := 2
		actual := longestOnes(nums, k)
		if actual != expect {
			t.Errorf("expect: %v, actual: %v", expect, actual)
		}
	})

	t.Run("simple test", func(t *testing.T) {
		nums := []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}
		k := 3
		expect := 10
		actual := longestOnes(nums, k)
		if actual != expect {
			t.Errorf("expect: %v, actual: %v", expect, actual)
		}
	})

	t.Run("simple test", func(t *testing.T) {
		nums := []int{1, 1, 1, 1, 1, 1, 1}
		k := 0
		expect := 7
		actual := longestOnes(nums, k)
		if actual != expect {
			t.Errorf("expect: %v, actual: %v", expect, actual)
		}
	})

	t.Run("simple test", func(t *testing.T) {
		nums := []int{0, 0, 0, 0, 0, 0, 0}
		k := 100
		expect := 7
		actual := longestOnes(nums, k)
		if actual != expect {
			t.Errorf("expect: %v, actual: %v", expect, actual)
		}
	})
}
