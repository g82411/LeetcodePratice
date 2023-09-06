package _86_Interval_List_Intersections

func intervalIntersection(A [][]int, B [][]int) [][]int {
	i := 0
	j := 0
	var ans [][]int
	for i < len(A) && j < len(B) {
		a := A[i]
		b := B[j]
		if a[1] < b[0] {
			i++
		} else if a[0] > b[1] {
			j++
		} else if a[1] == b[0] {
			ans = append(ans, []int{a[1], b[0]})
			i++
		} else if b[1] == a[0] {
			ans = append(ans, []int{a[0], b[1]})
			j++
		} else {
			left := max(a[0], b[0])
			right := min(a[1], b[1])
			ans = append(ans, []int{left, right})
			if a[1] < b[1] {
				i++
			} else {
				j++
			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
