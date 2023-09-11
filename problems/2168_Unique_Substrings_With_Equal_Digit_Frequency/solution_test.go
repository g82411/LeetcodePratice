package _168_Unique_Substrings_With_Equal_Digit_Frequency

import "testing"

func TestEqualDigitFrequency(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		expect := 5
		actual := equalDigitFrequency("1212")
		if actual != expect {
			t.Errorf("except: %v, actual: %v", expect, actual)
		}
	})

	t.Run("Case 1", func(t *testing.T) {
		expect := 9
		actual := equalDigitFrequency("12321")
		if actual != expect {
			t.Errorf("except: %v, actual: %v", expect, actual)
		}
	})
}
