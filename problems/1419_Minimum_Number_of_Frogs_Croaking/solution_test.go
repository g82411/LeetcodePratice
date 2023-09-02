package _419_Minimum_Number_of_Frogs_Croaking

import "testing"

func TestMinimumNumberOfFrogsCroaking(t *testing.T) {
	testcases := []struct {
		croakOfFrogs string
		except       int
	}{
		{
			"croakcroak",
			1,
		},
		{
			"crcoakroak",
			2,
		},
		{
			"croakcrook",
			-1,
		},
		{
			"croakcroa",
			-1,
		},
		{
			"aoocrrackk",
			-1,
		},
	}
	for _, testcase := range testcases {
		actual := minNumberOfFrogs(testcase.croakOfFrogs)
		if actual != testcase.except {
			t.Errorf("croakOfFrogs: %v, except: %v, actual: %v", testcase.croakOfFrogs, testcase.except, actual)
		}
	}
}
