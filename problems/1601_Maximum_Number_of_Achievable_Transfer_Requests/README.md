# 1601. Maximum Number of Achievable Transfer Requests

注意看數值範圍，基本上可以用狀態窮舉
然後，注意這提要找到滿足最多的state，所以就是1最多的state

所以有一個Gosper's Hack

```go
package util
func GosperHack(n, k int) []int {
	var res []int
	state := (1 << k) - 1
	for state < 1 << n {
		res = append(res, state)
		c := state & - state
		r := state + c
		state = (((r ^ state) >> 2) / c) | r
    }
	return res
}
```
可以求出長度為n的bit中 有k個1的所有組合

