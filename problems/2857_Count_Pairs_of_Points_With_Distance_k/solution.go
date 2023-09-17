package _857_Count_Pairs_of_Points_With_Distance_k

// 注意xor的性質
// k = a + b
// a = x1 ^ x2 => x2 = x1 ^ a
// y1 ^ y2 = b => y2 = y1 ^ b
// hash table

func countPairs(coordinates [][]int, k int) int {
	var encodeCoord func(coord []int) int64
	encodeCoord = func(coord []int) int64 {
		x, y := coord[0], coord[1]
		return int64(x)*int64(1e6) + int64(y)
	}
	ret := 0
	for a := 0; a <= k; a++ {
		coordTable := make(map[int64]int)
		for _, coord := range coordinates {
			x1, y1 := coord[0], coord[1]
			x2 := a ^ x1
			y2 := (k - a) ^ y1
			code := encodeCoord([]int{x2, y2})
			ret += coordTable[code]
			newCode := encodeCoord([]int{x1, y1})
			coordTable[newCode]++
		}
	}
	return ret

}
