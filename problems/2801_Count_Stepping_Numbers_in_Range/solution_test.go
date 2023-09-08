package _801_Count_Stepping_Numbers_in_Range

import "testing"

func TestCountSteppingNumbers(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		expect := 10
		actual := countSteppingNumbers("1", "11")
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		expect := 2
		actual := countSteppingNumbers("90", "101")
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})

	t.Run("Case 3", func(t *testing.T) {
		expect := 7
		actual := countSteppingNumbers("44", "86")
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})

	t.Run("Case 4", func(t *testing.T) {
		expect := 18
		actual := countSteppingNumbers("21", "145")
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})

	t.Run("Overflow ", func(t *testing.T) {
		expect := 705298990
		actual := countSteppingNumbers("999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", "9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999")
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
}
