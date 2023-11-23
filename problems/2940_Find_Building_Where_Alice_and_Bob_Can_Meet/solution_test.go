package _940_Find_Building_Where_Alice_and_Bob_Can_Meet

// in query[i] {a, b} 找到第一個heights[k] > max(heights[a], heights[b])
// 所以從最右邊開始處理 然後把處理過的heights 放到有序容器裡面

import (
	"testing"
)

func TestLeftmostBuildingQueries(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		got := leftmostBuildingQueries([]int{5, 3, 8, 2, 6, 1, 4, 6}, [][]int{{0, 7}, {3, 5}, {5, 2}, {3, 0}, {1, 6}})
		want := []int{7, 6, -1, 4, 6}
		if !equal(got, want) {
			t.Errorf("got %v want %v given", got, want)
		}
	})

}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}
