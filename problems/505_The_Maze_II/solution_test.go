package _05_The_Maze_II

import "testing"

func TestShortestDistance(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		maze := [][]int{
			{0, 0, 1, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 1, 0},
			{1, 1, 0, 1, 1},
			{0, 0, 0, 0, 0},
		}
		start := []int{0, 4}
		destination := []int{4, 4}
		expected := 12
		result := shortestDistance(maze, start, destination)
		if result != expected {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})
	t.Run("test2", func(t *testing.T) {
		maze := [][]int{
			{0, 0, 1, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 1, 0, 1, 0},
			{1, 1, 0, 1, 1},
			{0, 0, 0, 0, 0},
		}
		start := []int{0, 4}
		destination := []int{3, 2}
		expected := -1
		result := shortestDistance(maze, start, destination)
		if result != expected {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})
}
