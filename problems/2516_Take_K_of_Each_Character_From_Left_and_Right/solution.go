package _516_Take_K_of_Each_Character_From_Left_and_Right

func takeCharacters(s string, k int) int {
	mostA := -k
	mostB := -k
	mostC := -k
	for _, c := range s {
		if c == 'a' {
			mostA++
		}
		if c == 'b' {
			mostB++
		}
		if c == 'c' {
			mostC++
		}
	}
	if !(mostA >= 0 && mostB >= 0 && mostC >= 0) {
		return -1
	}
	maxLength := 0
	r := 0
	aInWindow := 0
	bInWindow := 0
	cInWindow := 0
	for l := 0; l < len(s); l++ {
		for aInWindow <= mostA && bInWindow <= mostB && cInWindow <= mostC {
			maxLength = max(maxLength, r-l)
			if r == len(s) {
				break
			}
			c := s[r]
			if c == 'a' {
				aInWindow++
			}
			if c == 'b' {
				bInWindow++
			}
			if c == 'c' {
				cInWindow++
			}
			r++
		}
		c := s[l]
		if c == 'a' {
			aInWindow--
		}
		if c == 'b' {
			bInWindow--
		}
		if c == 'c' {
			cInWindow--
		}
	}
	return len(s) - maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
