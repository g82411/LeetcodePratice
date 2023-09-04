package _727_Largest_Submatrix_With_Rearrangements

import "testing"

func TestLargestSubmatrix(t *testing.T) {
	testcases := []struct {
		matrix [][]int
		expect int
	}{
		{
			[][]int{{0, 0, 1}, {1, 1, 1}, {1, 0, 1}},
			4,
		},
		{
			[][]int{{1, 0, 1, 0, 1}},
			3,
		},
		{
			[][]int{{1, 1, 0}, {1, 0, 1}},
			2,
		},
		{
			[][]int{{0, 0}, {0, 0}},
			0,
		},
	}
	for _, testcase := range testcases {
		actual := largestSubmatrix(testcase.matrix)
		if actual != testcase.expect {
			t.Errorf("matrix: %v, expect: %v, actual: %v", testcase.matrix, testcase.expect, actual)
		}
	}
}
