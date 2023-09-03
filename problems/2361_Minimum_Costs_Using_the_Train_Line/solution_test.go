package _361_Minimum_Costs_Using_the_Train_Line

import "testing"

func TestMinimumCosts(t *testing.T) {
	t.Run("Basic testcase", func(t *testing.T) {
		regular := []int{1, 6, 9, 5}
		express := []int{5, 2, 3, 10}
		expressCost := 8
		expect := []int64{1, 7, 14, 19}
		actual := minimumCosts(regular, express, expressCost)
		if len(actual) != len(expect) {
			t.Errorf("except: %v, actual: %v", expect, actual)
		}
		for i := range expect {
			if actual[i] != expect[i] {
				t.Errorf("except: %v, actual: %v", expect, actual)
			}
		}
	})

	t.Run("Basic testcase", func(t *testing.T) {
		regular := []int{11, 5, 13}
		express := []int{7, 10, 6}
		expressCost := 3
		expect := []int64{10, 15, 24}
		actual := minimumCosts(regular, express, expressCost)
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
