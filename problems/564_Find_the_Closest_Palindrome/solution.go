package _64_Find_the_Closest_Palindrome

import "math/big"

// 最近的candidate是
// s = ABCDE
// 那至少有一個很相近的ABCBA
// 那問題是有沒有反例勒？
// s = "19200" -> s' = "19291"
// 顯然有一個更接近的19191
// 所以我們要找到一大一小兩個接近的 去比較哪個更接近
func nearestPalindromic(n string) string {
	greaterString := makeGreaterPalindromic(&n)
	smallerString := makeSmallerPalindromic(&n)
	bigger, _ := new(big.Int).SetString(greaterString, 10)
	smaller, _ := new(big.Int).SetString(smallerString, 10)
	origin, _ := new(big.Int).SetString(n, 10)
	biggerDiff := new(big.Int).Sub(bigger, origin)
	smallerDiff := new(big.Int).Sub(origin, smaller)

	cmp := biggerDiff.Cmp(smallerDiff)
	if cmp == -1 {
		return greaterString
	}
	return smallerString
}

func makeGreaterPalindromic(n *string) string {
	s := []byte(*n)
	l := len(s)

	left, right := 0, (l - 1)
	for right >= left {
		s[right] = s[left]
		right--
		left++
	}
	// 如果大於 快樂return
	if string(s) > *n {
		return string(s)
	}
	// 如果小於 那就要來找條大魚
	// 中位 + 1 其他為反轉
	carry := 1
	for i := (l - 1) / 2; i >= 0; i-- {
		d := int(s[i]-'0') + carry
		if d <= 9 {
			s[i] = byte('0' + d)
			carry = 0
		} else {
			s[i] = byte('0')
			carry = 1
		}
		s[l-1-i] = s[i]
	}
	// 特例：999 -> 1001
	if carry == 1 {
		ret := make([]byte, l+1)
		for i := range ret {
			if i == 0 || i == l {
				ret[i] = '1'
				continue
			}
			ret[i] = '0'
		}
		return string(ret)
	}
	return string(s)
}

func makeSmallerPalindromic(n *string) string {
	s := []byte(*n)
	l := len(s)
	left, right := 0, (l - 1)
	for right >= left {
		s[right] = s[left]
		right--
		left++
	}
	if string(s) < *n {
		return string(s)
	}
	carry := 1
	for i := (l - 1) / 2; i >= 0; i-- {
		d := int(s[i]-'0') - carry
		if d >= 0 {
			s[i] = byte('0' + d)
			carry = 0
		} else {
			s[i] = '9'
			carry = 1
		}
		s[l-1-i] = s[i]
	}
	// 特例1000 -> 999
	if s[0] == '0' && l > 1 {
		ret := make([]byte, l-1)
		for i := range ret {
			ret[i] = '9'
		}
		return string(ret)
	}
	return string(s)
}
