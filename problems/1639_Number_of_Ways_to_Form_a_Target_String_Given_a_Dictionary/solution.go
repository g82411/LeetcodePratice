package _639_Number_of_Ways_to_Form_a_Target_String_Given_a_Dictionary

func numWays(words []string, target string) int {
	var dp [1005][1005]int64
	var count [1005][27]int
	const Mod = int(1e9 + 7)
	// dp[i][k] => how many ways to form target[0:i] using [0:k]
	// dp[i][k] = dp[i][k-1]
	// if can use word[k] to form target[i]
	// dp[i][k] += dp[i-1][k-1] * count(how many char word[k] == target[i])
	n := len(target)
	m := len(words[0])
	for _, word := range words {
		for i := range word {
			count[i+1][int(word[i]-'a')]++
		}
	}

	target = "#" + target
	// there only 1 ways to form target[0:0] using word[0:k]
	for k := 0; k <= m; k++ {
		dp[0][k] = 1
	}
	for i := 1; i <= n; i++ {
		for k := 1; k <= m; k++ {
			dp[i][k] = dp[i][k-1]
			dp[i][k] += dp[i-1][k-1] * int64(count[k][int(target[i]-'a')]) % int64(Mod)
			dp[i][k] %= int64(Mod)
		}
	}
	return int(dp[n][m] % int64(Mod))
}
