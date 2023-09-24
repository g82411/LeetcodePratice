package _99_Champagne_Tower

import "testing"

func TestChampagneTower(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		expect := 0.0
		actual := champagneTower(1, 1, 1)
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		expect := 0.5
		actual := champagneTower(2, 1, 1)
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
	t.Run("Case 3", func(t *testing.T) {
		expect := 1.0
		actual := champagneTower(100000009, 33, 17)
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
}
