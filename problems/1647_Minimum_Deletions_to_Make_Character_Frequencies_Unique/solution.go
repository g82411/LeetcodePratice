package _647_Minimum_Deletions_to_Make_Character_Frequencies_Unique

func minDeletions(s string) int {
	wordFrequence := make(map[rune]int)
	for _, c := range s {
		wordFrequence[c]++
	}
	result := 0
	duplicateFrequency := make(map[int]int)
	for _, value := range wordFrequence {
		f := value
		for duplicateFrequency[f] > 0 && f > 0 {
			f--
			result++
		}
		duplicateFrequency[f]++
	}
	return result
}
