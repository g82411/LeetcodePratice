package _6_Minimum_Window_Substring

import "testing"

func TestMinWindow(t *testing.T) {
	t.Run("Case1", func(t *testing.T) {
		expect := "BANC"
		actual := minWindow("ADOBECODEBANC", "ABC")
		if expect != actual {
			t.Errorf("except: %v, actual: %v", expect, actual)

		}
	})

	t.Run("Case1", func(t *testing.T) {
		expect := "a"
		actual := minWindow("a", "a")
		if expect != actual {
			t.Errorf("except: %v, actual: %v", expect, actual)

		}
	})

	t.Run("Case1", func(t *testing.T) {
		expect := ""
		actual := minWindow("a", "aa")
		if expect != actual {
			t.Errorf("except: %v, actual: %v", expect, actual)

		}
	})
	t.Run("Case1", func(t *testing.T) {
		expect := "ccabaasdghjasdghajsbbasdw"
		actual := minWindow("abbbbbbbbbbbbbbbbbbbbbbbbbacccabaasdghjasdghajsbbasdwwasdbbbb", "bbaccw")
		if expect != actual {
			t.Errorf("except: %v, actual: %v", expect, actual)

		}
	})
}
