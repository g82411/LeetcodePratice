package _827_Number_of_Beautiful_Integers_in_the_Range

import "testing"

func TestNumberOfBeautifulIntegers(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		expect := 2
		actual := numberOfBeautifulIntegers(10, 20, 3)
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})

	t.Run("Case 1", func(t *testing.T) {
		expect := 1
		actual := numberOfBeautifulIntegers(1, 10, 1)
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})

	t.Run("Case 1", func(t *testing.T) {
		expect := 0
		actual := numberOfBeautifulIntegers(5, 5, 2)
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})

}
