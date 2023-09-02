package _34_Gas_Statio

import "testing"

func TestCanCompleteCircular(t *testing.T) {
	testcases := []struct {
		gas    []int
		cost   []int
		except int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			[]int{3, 4, 5, 1, 2},
			3,
		},
		{
			[]int{2, 3, 4},
			[]int{3, 4, 3},
			-1,
		},
	}
	for _, testcase := range testcases {
		actual := canCompleteCircuit(testcase.gas, testcase.cost)
		if actual != testcase.except {
			t.Errorf("gas: %v, cost: %v, except: %v, actual: %v", testcase.gas, testcase.cost, testcase.except, actual)
		}
	}
}
