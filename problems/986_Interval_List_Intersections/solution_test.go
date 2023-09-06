package _86_Interval_List_Intersections

import "testing"

func TestIntervalIntersection(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		A := [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}
		B := [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}
		actual := intervalIntersection(A, B)
		expect := [][]int{{1, 2}, {5, 5}, {8, 10}, {15, 23}, {24, 24}, {25, 25}}
		if len(actual) != len(expect) {
			t.Errorf("except: %v, actual: %v", expect, actual)
		}
		for i := range expect {
			if len(actual[i]) != len(expect[i]) {
				t.Errorf("except: %v, actual: %v", expect, actual)
			}
			for j := range expect[i] {
				if actual[i][j] != expect[i][j] {
					t.Errorf("except: %v, actual: %v", expect, actual)
				}
			}
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		A := [][]int{{3, 5}, {9, 20}}
		B := [][]int{{4, 5}, {7, 10}, {11, 12}, {14, 15}, {16, 20}}
		actual := intervalIntersection(A, B)
		expect := [][]int{{4, 5}, {9, 10}, {11, 12}, {14, 15}, {16, 20}}
		if len(actual) != len(expect) {
			t.Errorf("except: %v, actual: %v", expect, actual)
		}
		for i := range expect {
			if len(actual[i]) != len(expect[i]) {
				t.Errorf("except: %v, actual: %v", expect, actual)
			}
			for j := range expect[i] {
				if actual[i][j] != expect[i][j] {
					t.Errorf("except: %v, actual: %v", expect, actual)
				}
			}
		}
	})
}
