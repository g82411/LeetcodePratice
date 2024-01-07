package _808_Soup_Servings

func soupServings(n int) float64 {
	var cache [5000][5000]float64
	if n >= 5000 {
		return 1.0
	}
	var dfs func(a, b int) float64
	dfs = func(a, b int) float64 {
		if a > 0 && b <= 0 {
			return 0.0
		}
		if a <= 0 && b > 0 {
			return 1.0
		}

		if a <= 0 && b <= 0 {
			return 0.5
		}

		if cache[a][b] > 0.0 {
			return cache[a][b]
		}
		cache[a][b] = 0.25 * (dfs(a-100, b) + dfs(a-75, b-25) + dfs(a-50, b-50) + dfs(a-25, b-75))
		return cache[a][b]
	}
	// if m >= 5000, the prob will nearly to 1
	m := min(5000, n)
	return dfs(m, m)
}
