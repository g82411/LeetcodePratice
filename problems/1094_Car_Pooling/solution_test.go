package _094_Car_Pooling

import "testing"

func TestCarPooling(t *testing.T) {
	testcases := []struct {
		trips    [][]int
		capacity int
		expect   bool
	}{
		{
			[][]int{{2, 1, 5}, {3, 3, 7}},
			4,
			false,
		},
		{
			[][]int{{2, 1, 5}, {3, 3, 7}},
			5,
			true,
		},
	}
	for _, testcase := range testcases {
		actual := carPooling(testcase.trips, testcase.capacity)
		if actual != testcase.expect {
			t.Errorf("trips: %v, capacity: %v, expect: %v, actual: %v", testcase.trips, testcase.capacity, testcase.expect, actual)
		}
	}
}
