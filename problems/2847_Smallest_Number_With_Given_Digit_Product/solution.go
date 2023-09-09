package _847_Smallest_Number_With_Given_Digit_Product

import "strconv"

// TODO: 多練習
// 這題比較難想 為什麼只能是2~9 因為10^1 * 10^1 = 10^2 一共三位
// 但是10^1 有兩位 兩位 + 兩位 一共有四位
func smallestNumber(n int64) string {
	if n == 1 {
		return "1"
	}
	ret := ""
	for i := 9; i > 1; i-- {
		for n%int64(i) == 0 {
			n /= int64(i)
			ret += strconv.Itoa(i)
		}
	}
	inverse_ret := ""
	for i := len(ret) - 1; i >= 0; i-- {
		inverse_ret += string(ret[i])
	}
	if n != 1 {
		return "-1"
	}
	return inverse_ret
}
