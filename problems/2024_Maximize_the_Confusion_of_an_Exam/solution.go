package _024_Maximize_the_Confusion_of_an_Exam

func maxConsecutiveAnswers(answerKey string, k int) int {
	chars := []rune{'T', 'F'}
	maxLength := 0
	for _, char := range chars {
		l := 0
		r := 0
		requireChange := 0
		for _, ans := range answerKey {
			if ans != char {
				requireChange++
			}
			for requireChange > k {
				if rune(answerKey[l]) != char {
					requireChange--
				}
				l++
			}
			maxLength = max(maxLength, r-l+1)
			r++
		}
	}
	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
