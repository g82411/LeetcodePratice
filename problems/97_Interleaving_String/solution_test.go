package _97_Interleaving_String

import "testing"

func TestIsInterleave(t *testing.T) {
	tests := []struct {
		s1   string
		s2   string
		s3   string
		want bool
	}{
		{"aabcc", "dbbca", "aadbbcbcac", true},
		{"aabcc", "dbbca", "aadbbbaccc", false},
		{"", "", "", true},
	}
	for _, test := range tests {
		if got := isInterleave(test.s1, test.s2, test.s3); got != test.want {
			t.Errorf("isInterleave(%q, %q, %q) = %v; want %v\n", test.s1, test.s2, test.s3, got, test.want)
		}
	}
}
