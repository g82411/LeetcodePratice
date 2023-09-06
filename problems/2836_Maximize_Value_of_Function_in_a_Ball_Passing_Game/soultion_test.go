package _836_Maximize_Value_of_Function_in_a_Ball_Passing_Game

import "testing"

func TestGetMaxFunctionValue(t *testing.T) {
	t.Run("Basic test", func(t *testing.T) {
		expect := int64(6)
		actual := getMaxFunctionValue([]int{2, 0, 1}, 4)
		if actual != expect {
			t.Errorf("expect: %v, actual: %v", expect, actual)
		}
	})
	t.Run("Basic test", func(t *testing.T) {
		expect := int64(10)
		actual := getMaxFunctionValue([]int{1, 1, 1, 2, 3}, 3)
		if actual != expect {
			t.Errorf("expect: %v, actual: %v", expect, actual)
		}
	})
	t.Run("TLE test", func(t *testing.T) {
		expect := int64(1463)
		actual := getMaxFunctionValue([]int{14, 7, 18, 28, 12, 38, 14, 5, 22, 18, 19, 15, 24, 40, 0, 4, 21, 13, 18, 18, 3, 3, 2, 41, 14, 19, 15, 35, 10, 13, 28, 1, 21, 1, 29, 27, 17, 27, 26, 34, 0, 4}, 46)
		if actual != expect {
			t.Errorf("expect: %v, actual: %v", expect, actual)
		}
	})
}
