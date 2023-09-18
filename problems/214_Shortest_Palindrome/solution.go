package _14_Shortest_Palindrome

// 特定優化法
// 利用回文特性若 一個字串為回文字串
// 則 x := 0
// x ^ s[0] ^ s[1] ... ^ s[n-1] = 0 若len(s)為偶數
// x ^ s[0] ^ s[1] ... ^ s[n-1] = s[1/2] 若len(s)為奇數

func shortestPalindrome(s string) string {
	x := byte(0)
	n := len(s)
	for i := range []byte(s) {
		x ^= s[i]
	}
	isPalindrome := func(s string) bool {
		for i := 0; i < len(s)/2; i++ {
			if s[i] != s[len(s)-1-i] {
				return false
			}
		}
		return true
	}
	reserve := func(s string) string {
		res := ""
		for i := len(s) - 1; i >= 0; i-- {
			res += string(s[i])
		}
		return res
	}
	for i := n; i > 0; i-- {
		if x == 0 && i&1 == 0 && isPalindrome(s[0:i]) {
			return reserve(s[i:]) + s
		}
		if x == s[i/2] && i&1 == 1 && isPalindrome(s[0:i]) {
			return reserve(s[i:]) + s
		}
		x ^= s[i-1]
	}
	return ""
}

// 首先這題一定有解
// 假設有一個字串xyzk 那有一個必定解為kzyxyzk
// 也就是說一個回文都找不到 就把第二個字到最後一個字反轉
// 那問題就變成找到左邊最長的回文
//
//	func shortestPalindrome(s string) string {
//		padString := func() string {
//			t := "#"
//			for _, c := range s {
//				t += string(c)
//				t += "#"
//			}
//			return t
//		}
//		paddingString := padString()
//		maxCenter := -1
//		maxRight := -1
//		mockLength := len(paddingString)
//		p := make([]int, mockLength)
//		mostLeftRaidus := 1
//		for i := range paddingString {
//			radius := 0
//			if i < maxRight {
//				j := 2*maxCenter - i
//				radius = min(p[j], maxRight-i)
//			}
//			for i+radius < mockLength && i-radius >= 0 && (paddingString[i-radius] == paddingString[i+radius]) {
//				radius++
//			}
//			p[i] = radius - 1
//			if i+p[i] > maxRight {
//				maxRight = i + p[i]
//				maxCenter = i
//			}
//			if i-p[i] == 0 {
//				mostLeftRaidus = p[i]
//			}
//		}
//		resultPadding := s[mostLeftRaidus:]
//		// res := ""
//		// for i := len(resultPadding) - 1; i >= 0; i-- {
//		//     res += string(resultPadding[i])
//		// }
//		return reverse(resultPadding) + s
//	}
//
//	func min(a, b int) int {
//		if a < b {
//			return a
//		}
//		return b
//	}
//func reverse(s string) string {
//	bs := []byte(s)
//	for i, j := 0, len(bs)-1; i < j; i, j = i+1, j-1 {
//		bs[i], bs[j] = bs[j], bs[i]
//	}
//
//	return string(bs)
//}
