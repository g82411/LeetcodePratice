package _790_Maximum_Number_of_Groups_With_Increasing_Length

import "sort"

// 想了好久
// 首先簡單的想這題可以看成這樣
// 假設ul = [1,2,3,4,5] (先排序)
// 則可以構成 轉90度就是題目要的東西
//             5
//          4  5
//       3  4  5
//    2  3  4  5
// 1  2  3  4  5
// 問題是缺角怎麼半
// 設ul = [1,2,4,4,4]
// 則
// //                 4
//                4   5
//           3    4   5
//      2    3    4   5
// 1    2    3    4   5
// 可以跟前面借一位
// 所以假設Bin search
// l = 0 r = n * (n + 1) / 2 => n = len(ul)
// 設t = m
// 則
// for u in ul {
//     欠 = 欠 + t - u
//     t--
// }
// return 欠 >= 0
// 注意有一種狀況是不允許的 就是不可以拿出現頻率高去補低
// 設ul = [1,1,1,1,9999]
// //                 5
//                5   5
//           5    5   5
//      5    5    5   5
// 1    2    3    4   5
// 這樣會有重複的數字, 所以要確保欠款 <= 0
// for u in ul {
//     欠 = 欠 + t - u
//     欠 = min(欠, 0)
//     t--
// }

// func maxIncreasingGroups(usageLimits []int) int {
//     sort.Slice(usageLimits, func(i, j int ) bool {
//         return usageLimits[i] > usageLimits[j]
//     })
//     check := func (t int) bool {
//         o := 0
//         for _, u := range usageLimits {
//             o = min(0, o + u - t)
//             if t > 0 {
//                 t--
//             }
//         }
//         return o == 0 && t == 0
//     }
//     n := len(usageLimits)
//     l := 0
//     r := n * (n+1) / 2
//     ans := 0
//     for l <= r {
//         m := l + (r - l) / 2
//         if check(m) {
//             ans = m
//             l = m + 1
//         } else {
//             r = m - 1
//         }
//     }
//     return ans
// }

// 延續小可以補大 但大部可以捕小的精神
// 只要由小到大往上加 當s[n] > 1/2 * j * j+1
// 我們就可以假設三角形存在
func maxIncreasingGroups(ul []int) int {
	sort.Ints(ul)
	ans := 0
	s := 0
	triNumber := 1
	for _, num := range ul {
		s += num
		if s >= triNumber {
			ans++
			triNumber += ans + 1
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
