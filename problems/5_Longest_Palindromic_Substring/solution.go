package __Longest_Palindromic_Substring

import "strings"

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	var padString func() string
	padString = func() string {
		t := "#"
		for i := range s {
			t += string(s[i])
			t += "#"
		}
		return t
	}
	padStr := padString()
	n := len(padStr)
	p := make([]int, n)
	maxRight := 0
	maxCenter := 0
	for i := 0; i < n; i++ {
		radius := 0
		if i <= maxRight {
			j := maxCenter*2 - i
			radius = min(p[j], maxRight-i)
			for i+radius < n && i-radius >= 0 && padStr[i+radius] == padStr[i-radius] {
				radius++
			}
		} else {
			for i+radius < n && i-radius >= 0 && padStr[i+radius] == padStr[i-radius] {
				radius++
			}
		}
		p[i] = radius - 1
		if p[i]+i > maxRight {
			maxRight = i + p[i]
			maxCenter = i
		}
	}
	maxRadius := 0
	index := -1
	for i, radius := range p {
		if maxRadius < radius {
			index = i
			maxRadius = radius
		}
	}
	result := padStr[index-maxRadius : index+maxRadius]
	result = strings.Replace(result, "#", "", -1)

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 雙指針算法，稍微要記一下
//func longestPalindrome(s string) string {
//	if len(s) == 0 {
//		return ""
//	}
//
//	var left, right, previousLeft, previousRight int
//	for right < len(s) {
//		for right+1 < len(s) && s[left] == s[right+1] {
//			right++
//		}
//		for left-1 >= 0 && right+1 < len(s) && s[left-1] == s[right+1] {
//			left--
//			right++
//		}
//		if right-left > previousRight-previousLeft {
//			previousRight = right
//			previousLeft = left
//		}
//
//		left = (left+right)/2 + 1
//		right = left
//	}
//	return s[previousLeft : previousRight+1]
//}
