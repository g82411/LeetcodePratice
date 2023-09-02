package _90_Domino_and_Tromino_Tiling

import "testing"

func TestNumTilings(t *testing.T) {
	var cases = []struct {
		n      int
		output int
	}{
		{
			n:      3,
			output: 5,
		},
		{
			n:      4,
			output: 11,
		},
	}
	for _, c := range cases {
		x := numTilings(c.n)
		if x != c.output {
			t.Errorf("expected: %v, actual: %v", c.output, x)
		}
	}
}
