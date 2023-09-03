package _171_Removing_Minimum_Number_of_Magic_Beans

import "testing"

func TestNinimumRemoval(t *testing.T) {
	t.Run("Default case 1", func(t *testing.T) {
		beans := []int{4, 1, 6, 5}
		actual := minimumRemoval(beans)
		expected := int64(4)
		if actual != expected {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})

	t.Run("Default case 2", func(t *testing.T) {
		beans := []int{2, 10, 3, 2}
		actual := minimumRemoval(beans)
		expected := int64(7)
		if actual != expected {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	})
}
