package _168_Unique_Substrings_With_Equal_Digit_Frequency

// 問題是直接做hash他不香嗎？
// 直接做hash的版本在下面，問題在於casting & 字串相加會變得很慢
// rolling 主要功能是把比較變成數字比較，其結果會快很多
func equalDigitFrequency(s string) int {
	ans := make(map[int64]bool)
	var digitFrequency map[int]int
	mod := int64(1e9 + 7)
	n := len(s)
	for i := 0; i < n; i++ {
		digitFrequency = make(map[int]int)
		hash := int64(0)
		for j := i; j < n; j++ {
			charDigit := s[j] - '0'
			hash = hash * 11 % mod
			hash = hash + int64(charDigit+1)
			digitFrequency[int(charDigit)]++
			mostFrequency := -1
			hasSameDigit := true
			for _, value := range digitFrequency {
				if mostFrequency == -1 {
					mostFrequency = value
					continue
				}
				if mostFrequency != value {
					hasSameDigit = false
					break
				}
			}
			if hasSameDigit {
				ans[hash] = true
			}
		}
	}
	return len(ans)
}

//func equalDigitFrequency(s string) int {
//	ans := make(map[string]bool)
//	var digitFrequency map[int]int
//	n := len(s)
//	for i := 0; i < n; i++ {
//		digitFrequency = make(map[int]int)
//		hash := ""
//		for j := i; j < n; j++ {
//			hash += string(s[j])
//			digitFrequency[int(s[j]-'0')]++
//			mostFrequency := -1
//			hasSameDigit := true
//			for _, value := range digitFrequency {
//				if mostFrequency == -1 {
//					mostFrequency = value
//					continue
//				}
//				if mostFrequency != value {
//					hasSameDigit = false
//					break
//				}
//			}
//			if hasSameDigit {
//				ans[hash] = true
//			}
//		}
//	}
//	return len(ans)
//}
