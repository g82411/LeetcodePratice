package _52_Student_Attendance_Record_II

import "testing"

func TestCheckRecord(t *testing.T) {
	testcases := []struct {
		n      int
		except int
	}{
		{
			2,
			8,
		},
		{
			1,
			3,
		},
		{
			10101,
			183236316,
		},
	}
	for _, testcase := range testcases {
		actual := checkRecord(testcase.n)
		if actual != testcase.except {
			t.Errorf("n: %v, expect: %v, actual: %v", testcase.n, testcase.except, actual)
		}
	}
}
