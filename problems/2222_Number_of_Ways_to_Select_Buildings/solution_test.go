package _222_Number_of_Ways_to_Select_Buildings

import "testing"

func TestNumberOfWays(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		expect := int64(2)
		actual := numberOfWays("0101")
		if expect != actual {
			t.Errorf("Expect %v actual %v", expect, actual)
		}
	})
}
