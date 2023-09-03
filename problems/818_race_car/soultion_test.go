package _18_race_car

import "testing"

func TestRaceCar(t *testing.T) {
	testcases := []struct {
		target int
		expect int
	}{
		{
			3,
			2,
		},
		{
			6,
			5,
		},
		{
			5,
			7,
		},
		{
			4,
			5,
		},
	}
	for _, testcase := range testcases {
		actual := racecar(testcase.target)
		if actual != testcase.expect {
			t.Errorf("target: %v, expect: %v, actual: %v", testcase.target, testcase.expect, actual)
		}
	}
}
