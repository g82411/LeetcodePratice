package _47_Palindromic_Substrings

// consider P[i] as max expandable radius of Palindromic str which center is i
// mostRight
// most Center
// X X {X X X X X X X X X X X X X X X X X}X X X X X X
//
//	    i                 c             j mR
//	r = min{P[j], maxR - i}
//
// 馬拉車算法
func countSubstrings(s string) int {
	var paddingString func() string
	paddingString = func() string {
		t := "#"
		for _, c := range s {
			t += string(c)
			t += "#"
		}
		return t
	}
	strWithPadding := paddingString()
	maxCenter := -1
	maxRight := -1
	mockLength := len(strWithPadding)
	p := make([]int, mockLength)
	for i := range strWithPadding {
		radius := 0
		// i 在最大回文字串裡面，所以他對面的j也在最大回文字串
		if i <= maxRight {
			// 試想
			// 0 1 2 3 4 5 6
			// X X X X X X X
			//   j   C   i R
			// j = 3 * 2 - 5 = 1
			j := maxCenter*2 - i
			// 在MaxRight 內至少有可能是回文
			// 如果p[j] - j, p[j] + j 都是回文 則 p[i] - i, p[i] + i有可能是回文
			// 兩個取小往外推
			radius = min(p[j], maxRight-i)
			for i+radius < mockLength && i-radius >= 0 && (strWithPadding[i-radius] == strWithPadding[i+radius]) {
				radius++
			}
		} else {
			// i在無人之地 自己慢慢推
			for i+radius < mockLength && i-radius >= 0 && (strWithPadding[i-radius] == strWithPadding[i+radius]) {
				radius++
			}
		}
		p[i] = radius - 1
		if i+p[i] > maxRight {
			maxCenter = i
			maxRight = i + p[i]
		}
	}
	ret := 0
	for i := range p {
		ret += (p[i] + 1) / 2
	}
	return ret
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
