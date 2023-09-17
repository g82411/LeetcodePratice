package _852_Sum_of_Remoteness_of_All_Cells

type UnionFind struct {
	parent []int64
}

func NewUnionFind(n int64) UnionFind {
	parent := make([]int64, n)
	for i := range parent {
		parent[i] = int64(i)
	}
	return UnionFind{
		parent,
	}
}

func (u UnionFind) Find(x int64) int64 {
	if x != u.parent[x] {
		u.parent[x] = u.Find(u.parent[x])
	}
	return u.parent[x]
}

func (u UnionFind) Union(x, y int64) {
	rx := u.Find(x)
	ry := u.Find(y)
	if rx == ry {
		return
	}
	if rx > ry {
		u.parent[rx] = ry
		return
	}
	u.parent[ry] = rx
}

func sumRemoteness(grid [][]int) int64 {
	dir := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	n := len(grid)
	uf := NewUnionFind(int64(n) * int64(n))
	encodeCoord := func(x, y int) int64 {
		code := int64(y * n)
		code += int64(x)
		return code
	}
	visited := make([][]bool, n)
	rootSumMapping := make(map[int64]int64)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	var dfs func(x, y int)
	total := int64(0)
	dfs = func(x, y int) {
		if grid[y][x] == -1 {
			return
		}
		if visited[y][x] {
			return
		}
		visited[y][x] = true
		oriCoord := encodeCoord(x, y)
		total += int64(grid[y][x])
		for _, d := range dir {
			dy, dx := d[0], d[1]
			nx := x + dx
			ny := y + dy
			if !(nx >= 0 && ny >= 0 && nx < n && ny < n) {
				continue
			}
			if grid[ny][nx] == -1 {
				continue
			}
			nextCoord := encodeCoord(nx, ny)
			uf.Union(oriCoord, nextCoord)
			dfs(nx, ny)
		}
	}
	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			if grid[row][col] == -1 {
				continue
			}
			dfs(col, row)
			r := encodeCoord(col, row)
			rootSumMapping[uf.Find(r)] += int64(grid[row][col])
		}
	}
	var res int64
	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			if grid[row][col] == -1 {
				continue
			}
			root := uf.Find(encodeCoord(col, row))
			res += total - rootSumMapping[root]
		}
	}
	return res

}
