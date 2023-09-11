package _156_Find_Substring_With_Given_Hash_Value

import "testing"

func TestSubStrHash(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		expect := "ee"
		actual := subStrHash("leetcode", 7, 20, 2, 0)
		if actual != expect {
			t.Errorf("except: %v, actual: %v", expect, actual)
		}
	})

	t.Run("case 2", func(t *testing.T) {
		expect := "fbx"
		actual := subStrHash("fbxzaad", 31, 100, 3, 32)
		if actual != expect {
			t.Errorf("except: %v, actual: %v", expect, actual)
		}
	})

}
