package _99_The_Maze_III

import "testing"

func TestFindShortestWay(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		maze := [][]int{
			{0, 0, 0, 0, 0},
			{1, 1, 0, 0, 1},
			{0, 0, 0, 0, 0},
			{0, 1, 0, 0, 1},
			{0, 1, 0, 0, 0},
		}
		start := []int{4, 3}
		destination := []int{0, 1}
		got := findShortestWay(maze, start, destination)
		want := "lul"
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		maze := [][]int{
			{0, 0, 0, 0, 0},
			{1, 1, 0, 0, 1},
			{0, 0, 0, 0, 0},
			{0, 1, 0, 0, 1},
			{0, 1, 0, 0, 0},
		}
		start := []int{4, 3}
		destination := []int{3, 0}
		got := findShortestWay(maze, start, destination)
		want := "impossible"
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
	t.Run("Case 3", func(t *testing.T) {
		maze := [][]int{
			{0, 0, 0, 0, 0},
			{1, 1, 0, 0, 1},
			{0, 0, 0, 0, 0},
			{0, 1, 0, 0, 1},
			{0, 1, 0, 0, 0},
		}
		start := []int{4, 3}
		destination := []int{1, 4}
		got := findShortestWay(maze, start, destination)
		want := "urd"
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
}
