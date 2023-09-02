package _318_Number_of_Distinct_Roll_Sequences

import "testing"

func TestDistinctSequences(t *testing.T) {
	testcases := []struct {
		n      int
		except int
	}{
		{
			4,
			184,
		},
		{
			2,
			22,
		},
	}
	for _, testcase := range testcases {
		actual := distinctSequences(testcase.n)
		if actual != testcase.except {
			t.Errorf("n: %v, expect: %v, actual: %v", testcase.n, testcase.except, actual)
		}
	}
}
