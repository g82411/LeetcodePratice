package _842_Count_K_Subsequences_of_a_String_With_Maximum_Beauty

import "testing"

func TestCountKSubsequencesWithMaxBeauty(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		expect := 4
		actual := countKSubsequencesWithMaxBeauty("bcca", 2)
		if expect != actual {
			t.Errorf("expect: %v, actual: %v", expect, actual)
		}
	})

	t.Run("case 1", func(t *testing.T) {
		expect := 2
		actual := countKSubsequencesWithMaxBeauty("abbcd", 4)
		if expect != actual {
			t.Errorf("expect: %v, actual: %v", expect, actual)
		}
	})
}
