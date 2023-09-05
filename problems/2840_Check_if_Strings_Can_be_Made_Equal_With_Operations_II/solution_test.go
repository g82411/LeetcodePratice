package _840_Check_if_Strings_Can_be_Made_Equal_With_Operations_II

import "testing"

func TestCheckStrings(t *testing.T) {
	t.Run("Basic case", func(t *testing.T) {
		s1 := "abcdba"
		s2 := "cabdab"
		except := true
		actual := checkStrings(s1, s2)
		if actual != except {
			t.Errorf("s1: %v, s2: %v, except: %v, actual: %v", s1, s2, except, actual)
		}
	})

	t.Run("Basic case", func(t *testing.T) {
		s1 := "abe"
		s2 := "bea"
		except := false
		actual := checkStrings(s1, s2)
		if actual != except {
			t.Errorf("s1: %v, s2: %v, except: %v, actual: %v", s1, s2, except, actual)
		}
	})

}
