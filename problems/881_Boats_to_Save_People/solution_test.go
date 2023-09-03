package _81_Boats_to_Save_People

import "testing"

func TestNumRescueBoats(t *testing.T) {
	t.Run("Basic test case", func(t *testing.T) {
		actual := numRescueBoats([]int{1, 2}, 3)
		expect := 1
		if actual != expect {
			t.Errorf("except: %v, actual: %v", expect, actual)
		}
	})

	t.Run("Basic test case", func(t *testing.T) {
		actual := numRescueBoats([]int{3, 2, 2, 1}, 3)
		expect := 3
		if actual != expect {
			t.Errorf("except: %v, actual: %v", expect, actual)
		}
	})
}
