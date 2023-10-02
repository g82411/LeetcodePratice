package _872_Maximum_Number_of_K_Divisible_Components

func maxKDivisibleComponents(n int, edges [][]int, values []int, k int) int {
	var dfs func(curr int) int
	dp := make([]int, n)
	res := 0
	for i := 0; i < n; i++ {
		dp[i] = values[i]
	}
	neighbors := make([][]int, n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		neighbors[u] = append(neighbors[u], v)
		neighbors[v] = append(neighbors[v], u)
	}
	visited := make([]bool, n)
	dfs = func(curr int) int {
		visited[curr] = true
		for _, neighbor := range neighbors[curr] {
			if visited[neighbor] {
				continue
			}
			dp[curr] += dfs(neighbor)
		}
		if dp[curr]%k == 0 {
			res++
			return 0
		}
		return dp[curr]
	}
	dfs(0)
	return res

}
