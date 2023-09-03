package _516_Take_K_of_Each_Character_From_Left_and_Right

import "testing"

func TestTakeCharacters(t *testing.T) {
	t.Run("Basic case", func(t *testing.T) {
		s := "aabaaaacaabc"
		k := 2
		except := 8
		actual := takeCharacters(s, k)
		if actual != except {
			t.Errorf("s: %v, k: %v, except: %v, actual: %v", s, k, except, actual)
		}
	})
	t.Run("Basic case", func(t *testing.T) {
		s := "a"
		k := 1
		except := -1
		actual := takeCharacters(s, k)
		if actual != except {
			t.Errorf("s: %v, k: %v, except: %v, actual: %v", s, k, except, actual)
		}
	})
}
