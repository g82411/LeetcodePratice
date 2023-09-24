package _867_Count_Valid_Paths_in_a_Tree

// 筆記：
// 質數篩法O(n)
// 把質數找出來之後，用UnionFind處理，
// 找到prime的鄰居，所有的鄰居都可以跟prime形成不重複的一對

type UnionFind struct {
	parents []int
	ranks   []int
	size    []int
}

func NewUnionFind(n int) *UnionFind {
	parents := make([]int, n)
	ranks := make([]int, n)
	size := make([]int, n)
	for i := range ranks {
		parents[i] = i
		ranks[i] = i
		size[i] = 1
	}
	return &UnionFind{
		parents,
		ranks,
		size,
	}
}

func (u UnionFind) Find(x int) int {
	if u.parents[x] != x {
		u.parents[x] = u.Find(u.parents[x])
	}
	return u.parents[x]
}

func (u UnionFind) Union(x, y int) {
	rx := u.Find(x)
	ry := u.Find(y)
	if rx == ry {
		return
	}
	rankX := u.ranks[rx]
	rankY := u.ranks[ry]
	if rankX > rankY {
		u.parents[rx] = ry
		u.ranks[ry] += rx
		u.size[ry] += u.size[rx]
		return
	}
	u.parents[ry] = rx
	u.ranks[rx] += ry
	u.size[rx] += u.size[ry]
}

func countPaths(n int, edges [][]int) int64 {
	isPrime := make([]bool, n+1)
	uf := NewUnionFind(n + 1)
	for i := range isPrime {
		isPrime[i] = true
	}
	// 質數篩好像是O(n ^ 2) ?
	for i := 2; i <= n; i++ {
		if !isPrime[i] {
			continue
		}
		for k := 2; k <= n; k++ {
			if i*k > n {
				break
			}
			isPrime[i*k] = false
		}
	}
	isPrime[1] = false
	neighbor := make(map[int][]int)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		neighbor[u] = append(neighbor[u], v)
		neighbor[v] = append(neighbor[v], u)
		if !(isPrime[u] || isPrime[v]) {
			uf.Union(u, v)
			continue
		}
	}
	res := int64(0)
	for i := 1; i <= n; i++ {
		if !isPrime[i] {
			continue
		}
		sum := int64(1)
		var directNode []int
		for _, j := range neighbor[i] {
			if isPrime[j] {
				continue
			}
			directConnect := uf.size[uf.Find(j)]
			sum += int64(directConnect)
			directNode = append(directNode, directConnect)
		}
		// 這邊有點小trick
		// 考慮到一種可能
		//       2
		//      / \
		//     6   9
		//    /   /
		//   8   4
		// size 當然包含
		// 2-6
		// 2-9
		// 2-6-8
		// 2-9-4
		// 但有3種沒有考慮
		// 6-2-9
		// 8-6-2-9
		// 8-6-2-9-4
		// 所以考慮prime p的鄰居x, y, z Size a,b,c
		// 少考慮到的數量共 a * (b + c) + b * c
		for _, s := range directNode {
			sum -= int64(s)
			res += int64(s) * sum
		}
	}
	return res

}
