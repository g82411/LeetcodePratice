package _840_Check_if_Strings_Can_be_Made_Equal_With_Operations_II

func checkStrings(s1 string, s2 string) bool {
	charFrequency1Even := make([]byte, 26)
	charFrequency1Odd := make([]byte, 26)
	charFrequency2Even := make([]byte, 26)
	charFrequency2Odd := make([]byte, 26)

	for i := 0; i < len(s1); i++ {
		if i%2 == 0 {
			charFrequency1Even[int(s1[i]-'a')]++
			charFrequency2Even[int(s2[i]-'a')]++
			continue
		}
		charFrequency1Odd[int(s1[i]-'a')]++
		charFrequency2Odd[int(s2[i]-'a')]++
	}
	for i := 0; i < 26; i++ {
		if charFrequency1Even[i] != charFrequency2Even[i] {
			return false
		}
		if charFrequency1Odd[i] != charFrequency2Odd[i] {
			return false
		}
	}
	return true
}
