package _827_Number_of_Beautiful_Integers_in_the_Range

import "strconv"

func numberOfBeautifulIntegers(low int, high int, k int) int {
	// 就...2801改一改
	var determinFirstDigit func(num int) int
	var dfs func(remainLength, digitDiff, reminder int, isSame bool, strNum *string) int
	var memo [][22][22]map[bool]int

	dfs = func(remainLength, digitDiff, reminder int, isSame bool, strNum *string) int {
		n := len(*strNum)
		if remainLength == 0 {
			if digitDiff == 0 && reminder == 0 {
				return 1
			}
			return 0
		}
		if _, exist := memo[remainLength][digitDiff+20][reminder][isSame]; exist {
			return memo[remainLength][digitDiff+20][reminder][isSame]
		}
		ret := 0
		if !isSame {
			for d := 0; d <= 9; d++ {
				ret += dfs(remainLength-1, digitDiff+determinDiff(d), (reminder*10+d)%k, false, strNum)
			}
		} else {
			firstDigit := int((*strNum)[n-remainLength] - '0')
			for i := 0; i < firstDigit; i++ {
				ret += dfs(remainLength-1, digitDiff+determinDiff(i), (reminder*10+i)%k, false, strNum)
			}
			ret += dfs(remainLength-1, digitDiff+determinDiff(firstDigit), (reminder*10+firstDigit)%k, true, strNum)
		}
		memo[remainLength][digitDiff+20][reminder][isSame] = ret
		return ret
	}
	determinFirstDigit = func(num int) int {
		//	老規矩
		strNum := strconv.Itoa(num)

		memo = make([][22][22]map[bool]int, 11)
		for i := 0; i < 11; i++ {
			for j := 0; j < 22; j++ {
				for g := 0; g < 22; g++ {
					memo[i][j][g] = make(map[bool]int)
				}
			}

		}
		n := len(strNum)
		res := 0
		for i := 2; i < n; i += 2 {
			for d := 1; d <= 9; d++ {
				res += dfs(i-1, determinDiff(d), d%k, false, &strNum)
			}
		}
		if n%2 == 0 {
			firstDigit := int(strNum[0] - '0')
			for i := 1; i < firstDigit; i++ {
				res += dfs(n-1, determinDiff(i), i%k, false, &strNum)
			}
			res += dfs(n-1, determinDiff(firstDigit), firstDigit%k, true, &strNum)
		}
		return res
	}
	return determinFirstDigit(high) - determinFirstDigit(low-1)
}

func determinDiff(n int) int {
	if n%2 == 0 {
		return 1
	}
	return -1
}
