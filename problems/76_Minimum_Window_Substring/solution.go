package _6_Minimum_Window_Substring

import "math"

func minWindow(s string, t string) string {
	counter := 0
	lenS := len(s)
	lenT := len(t)
	tFrequency := make(map[byte]int)
	sFrequency := make(map[byte]int)
	shortestString := ""
	if lenS < lenT {
		return shortestString
	}
	for i := 0; i < len(t); i++ {
		tFrequency[t[i]]++
	}
	l, r := 0, 0
	minLength := math.MaxInt

	for r < lenS {
		for counter < lenT {
			if r >= lenS {
				break
			}
			sFrequency[s[r]]++
			if tFrequency[s[r]] >= sFrequency[s[r]] {
				counter++
			}
			r++
		}
		for counter == lenT {
			minL := r - l + 1
			if minL < minLength {
				minLength = minL
				shortestString = s[l:r]
			}
			sFrequency[s[l]]--
			if tFrequency[s[l]] > sFrequency[s[l]] {
				counter--
			}
			l++
		}
	}
	return shortestString
}
