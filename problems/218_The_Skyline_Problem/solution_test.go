package _18_The_Skyline_Problem

import "testing"

func TestGetSkyline(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		got := getSkyline([][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12},
			{15, 20, 10}, {19, 24, 8}})
		want := [][]int{{2, 10}, {3, 15}, {7, 12}, {12, 0},
			{15, 10}, {20, 8}, {24, 0}}
		if !equal(got, want) {
			t.Errorf("got %v want %v given", got, want)
		}
	})
}

func equal(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i][0] != v[0] || b[i][1] != v[1] {
			return false
		}
	}
	return true
}
