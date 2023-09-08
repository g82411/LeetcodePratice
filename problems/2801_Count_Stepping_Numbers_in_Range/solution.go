package _801_Count_Stepping_Numbers_in_Range

import "math"

// TODO: 感覺可以用dp做...而且TC不高
func countSteppingNumbers(low string, high string) int {
	// 關鍵點： 左邊界、右邊界、求邊界內之和
	// 要想到prefixSum(r) - prefixSum(l)的概念
	// 再來 要怎麼build prefixSum(r) 沒有思路就爆破
	// prefixSum(n) = prefixSum(len(n)) + [prefixSum(i) for i in range (n)]
	// 最後要注意low會被扣掉 要用check(low) 來確定要不要加回來 -> 注意前綴和的特性
	var findAllStepNumsLessThanNum func(num string) int64
	var dfs func(targetLen, prevDigit int, isAlignNum bool, maxNum *string) int64
	var check func(num string) int
	var memo [][10]map[bool]int64
	const mod = int64(1e9 + 7)
	dfs = func(targetLen, prevDigit int, isAlignNum bool, maxNum *string) int64 {
		if targetLen == 0 {
			return 1
		}
		if _, exist := memo[targetLen][prevDigit][isAlignNum]; exist {
			return memo[targetLen][prevDigit][isAlignNum]
		}
		n := len(*maxNum)
		ret := int64(0)
		// 如果前N位跟nums的前N位不align -> 可以隨意構造
		// e.g nums = "9234", prefixDigit = "234", "8234" "9233"...
		// 構造第一位為p + 1 & p - 1的數字和
		if !isAlignNum {
			if prevDigit+1 <= 9 {
				ret = (ret + dfs(targetLen-1, prevDigit+1, false, maxNum)) % mod
			}
			if prevDigit-1 >= 0 {
				ret = (ret + dfs(targetLen-1, prevDigit-1, false, maxNum)) % mod
			}
			ret %= mod
		} else {
			//如果align要注意最高位
			// e.g  nums "2345" "2343" ok "2356" not
			maxDigit := int((*maxNum)[n-targetLen] - '0')
			if prevDigit+1 < maxDigit {
				ret = (ret + dfs(targetLen-1, prevDigit+1, false, maxNum)) % mod
			} else if prevDigit+1 == maxDigit {
				ret = (ret + dfs(targetLen-1, prevDigit+1, true, maxNum)) % mod
			}
			ret %= mod
			if prevDigit-1 >= 0 && prevDigit-1 < maxDigit {
				ret = (ret + dfs(targetLen-1, prevDigit-1, false, maxNum)) % mod
			} else if prevDigit-1 >= 0 && prevDigit-1 == maxDigit {
				ret = (ret + dfs(targetLen-1, prevDigit-1, true, maxNum)) % mod
			}
			ret %= mod
		}
		memo[targetLen][prevDigit][isAlignNum] = ret
		return ret % mod
	}
	findAllStepNumsLessThanNum = func(num string) int64 {
		memo = make([][10]map[bool]int64, 105)
		for i := 0; i < 105; i++ {
			for j := 0; j < 10; j++ {
				memo[i][j] = make(map[bool]int64)
			}

		}

		n := len(num)
		ret := int64(0)
		// 構造length < n的所有數字
		for i := 1; i < n; i++ {
			for firstDigit := 1; firstDigit <= 9; firstDigit++ {
				ret = (ret + dfs(i-1, firstDigit, false, &num)) % mod
			}
		}
		firstDigitInNum := int(num[0] - '0')
		// 構造length == n 但ret[0] < nums[0]的數量
		for i := 1; i < firstDigitInNum; i++ {
			ret = (ret + dfs(n-1, i, false, &num)) % mod
		}
		// 構造length == n 但ret[0] == nums[0]的數量
		ret = (ret + dfs(n-1, firstDigitInNum, true, &num)) % mod
		return ret % mod
	}
	check = func(num string) int {
		for i := 1; i < len(num); i++ {
			// 注意string[i] type是uint8, 沒有處理就相減 會送你一個越位螺旋升天
			if int(math.Abs(float64(int(num[i])-int(num[i-1])))) != 1 {
				return 0
			}
		}
		return 1
	}
	ret := int64(0)
	ret += findAllStepNumsLessThanNum(high)
	ret -= findAllStepNumsLessThanNum(low)
	ret = (ret + mod) % mod
	ret = (ret + int64(check(low)) + mod) % mod
	return int(ret)
}
