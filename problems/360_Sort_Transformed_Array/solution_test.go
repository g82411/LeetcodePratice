package _60_Sort_Transformed_Array

import "testing"

func TestSortTransformedArray(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		actual := sortTransformedArray([]int{-4, -2, 2, 4}, 1, 3, 5)
		expect := []int{3, 9, 15, 33}
		if len(actual) != len(expect) {
			t.Errorf("except: %v, actual: %v", expect, actual)
		}
		for i := range expect {
			if actual[i] != expect[i] {
				t.Errorf("except: %v, actual: %v", expect, actual)
			}
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		actual := sortTransformedArray([]int{-4, -2, 2, 4}, -1, 3, 5)
		expect := []int{-23, -5, 1, 7}
		if len(actual) != len(expect) {
			t.Errorf("except: %v, actual: %v", expect, actual)
		}
		for i := range expect {
			if actual[i] != expect[i] {
				t.Errorf("except: %v, actual: %v", expect, actual)
			}
		}
	})

	t.Run("Case 3", func(t *testing.T) {
		actual := sortTransformedArray([]int{-4, -2, 2, 4}, 0, 3, 5)
		expect := []int{-7, -1, 11, 17}
		if len(actual) != len(expect) {
			t.Errorf("except: %v, actual: %v", expect, actual)
		}
		for i := range expect {
			if actual[i] != expect[i] {
				t.Errorf("except: %v, actual: %v", expect, actual)
			}
		}
	})

	t.Run("Case 3", func(t *testing.T) {
		actual := sortTransformedArray([]int{-4, -2, 2, 4}, 0, -3, 5)
		expect := []int{-7, -1, 11, 17}
		if len(actual) != len(expect) {
			t.Errorf("except: %v, actual: %v", expect, actual)
		}
		for i := range expect {
			if actual[i] != expect[i] {
				t.Errorf("except: %v, actual: %v", expect, actual)
			}
		}
	})
}
