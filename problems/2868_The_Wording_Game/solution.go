package _868_The_Wording_Game

func canAliceWin(a []string, b []string) bool {
	isCloselyGreater := func(w, z string) bool {
		if w <= z {
			return false
		}
		diff := w[0] - z[0]
		if diff < 0 {
			return false
		}
		if diff > 1 {
			return false
		}
		return true
	}
	turn := 1
	aWord := a[0]
	bWord := "123"
	for len(a) > 0 || len(b) > 0 {
		if turn%2 == 1 {
			for len(b) > 0 && !isCloselyGreater(b[0], aWord) {
				b = b[1:]
			}
			if len(b) > 0 && isCloselyGreater(b[0], aWord) {
				bWord = b[0]
				b = b[1:]
			} else {
				return true
			}

		} else {
			for len(a) > 0 && !isCloselyGreater(a[0], bWord) {
				a = a[1:]
			}
			if len(a) > 0 && isCloselyGreater(a[0], bWord) {
				aWord = a[0]
				a = a[1:]
			} else {
				return false
			}
		}
		turn++
	}
	return turn%2 == 1

}
