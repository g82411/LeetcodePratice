package _838_Maximum_Coins_Heroes_Can_Collect

import "testing"

func TestMaximumCoins(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		actual := maximumCoins([]int{1, 4, 2}, []int{1, 1, 5, 2, 3}, []int{2, 3, 4, 5, 6})
		expect := []int64{5, 16, 10}
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
		actual := maximumCoins([]int{5}, []int{2, 3, 1, 2}, []int{10, 6, 5, 2})
		expect := []int64{23}
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
		actual := maximumCoins([]int{4, 4}, []int{5, 7, 8}, []int{1, 1, 1})
		expect := []int64{0, 0}
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
		actual := maximumCoins([]int{8}, []int{13}, []int{15})
		expect := []int64{0}
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
