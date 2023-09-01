package _732_My_Calendar_III

import "testing"

func TestMyCalendarThree_Book(t *testing.T) {
	myCalendarThree := Constructor()
	testcases := []struct {
		startTime int
		endTime   int
		expect    int
	}{
		{
			10,
			20,
			1,
		},
		{
			50,
			60,
			1,
		},
		{
			10,
			40,
			2,
		},
		{
			5,
			15,
			3,
		},
	}
	for _, testcase := range testcases {
		actual := myCalendarThree.Book(testcase.startTime, testcase.endTime)
		if actual != testcase.expect {
			t.Errorf("startTime: %v, endTime: %v, expect: %v, actual: %v", testcase.startTime, testcase.endTime, testcase.expect, actual)
		}
	}

}
