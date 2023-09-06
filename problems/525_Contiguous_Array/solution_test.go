package _25_Contiguous_Array

import "testing"

func TestFindMaxLength(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		actual := findMaxLength([]int{0, 1})
		expect := 2
		if actual != expect {
			t.Errorf("except: %v, actual: %v", expect, actual)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		actual := findMaxLength([]int{0, 1, 0})
		expect := 2
		if actual != expect {
			t.Errorf("except: %v, actual: %v", expect, actual)
		}
	})
}
