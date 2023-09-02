package _01_Minimum_Swaps_To_Make_Sequences_Increasing

import "testing"

func TestMinSwap(t *testing.T) {
	t.Run("Default case 1", func(t *testing.T) {
		nums1 := []int{1, 3, 5, 4}
		nums2 := []int{1, 2, 3, 7}
		actual := minSwap(nums1, nums2)
		expected := 1
		if actual != expected {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})
	t.Run("Default case 2", func(t *testing.T) {
		nums1 := []int{0, 4, 4, 5, 9}
		nums2 := []int{0, 1, 6, 8, 10}
		actual := minSwap(nums1, nums2)
		expected := 1
		if actual != expected {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})
}
