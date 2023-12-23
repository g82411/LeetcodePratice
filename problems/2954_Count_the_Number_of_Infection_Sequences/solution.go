package _954_Count_the_Number_of_Infection_Sequences

// 1. 設定一個隨機序列
// X X X X S X X X S X X X S X X X
// X => no Sick, S => Sick
// 對於一個任一區間 S X X X S
// 有 以下幾種感染可能
// 1 2 3
// 3 2 1
// 1 3 2
// 3 1 2
// 可以推論 就是以下幾種 L R L, R R L, ...
// 注意 對於任意一種 決定 前兩個之後 剩下那個是確定的 沒有可能性
// 故對於一個被Sick包圍的飛Sick
// 有2 ^ (m - 1)種可能 m = len(no Sick)
// 2. 結果為序列 故結果最多有 n! (n = len(all no sick))
// 但4 我們已知其中任意被包覆的可能只有 2 ^ (m-1)
// 故總排列組合數為 n! / m! * 2 ^ (m - 1) / k! * 2 ^ (k - 1)
// 3. 注意結果需要取mod 但是除法並不保證同於定理 故需要找到逆元
// A  / b % M != A  % M / B % M
// 逆元性質:
// (A/B) % M = A % M * inv(B) % M
// 逆元求法:
// inv(x) => x ^ (M-2)

func numberOfSequence(n int, sick []int) int {
	const Mod = int64(1e9 + 7)
	fastPower := func(a int64) int64 {
		result := int64(1)
		base := a
		p := Mod - 2
		for p > 0 {
			// 如果當前的最低位是 1，則將當前的基數乘到結果中
			if p%2 == 1 {
				result = result * base % Mod
			}
			// 將基數平方
			base = base * base % Mod
			// 將指數右移一位
			p /= 2
		}
		return result
	}

	inv := func(a int64) int64 {
		return fastPower(a)
	}
	// 建立power & fact cache方便後續運算
	var power [100005]int64
	var fact [100005]int64
	power[0] = 1
	fact[0] = 1
	for i := 1; i <= n; i++ {
		power[i] = power[i-1] * 2 % Mod
		fact[i] = fact[i-1] * int64(i) % Mod
	}
	// 找到no sick 的gorup size
	var groups []int
	for i := range sick {
		if i == 0 {
			groups = append(groups, sick[i])
			continue
		}
		groups = append(groups, sick[i]-sick[i-1]-1)
	}
	//注意最後一組
	groups = append(groups, n-1-sick[len(sick)-1])

	totalSick := 0
	for _, g := range groups {
		totalSick += g
	}

	ret := fact[totalSick]
	// 同組的都先 / m!
	for _, size := range groups {
		// 注意這邊用了逆元 所以是*
		inverse := inv(fact[size])
		ret = ret * inverse % Mod
	}

	for i := 1; i < len(sick); i++ {
		size := sick[i] - sick[i-1] - 1
		if size > 0 {
			ret = ret * power[size-1] % Mod
		}
	}
	return int(ret)
}
