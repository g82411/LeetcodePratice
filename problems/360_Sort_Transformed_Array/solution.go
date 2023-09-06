package _60_Sort_Transformed_Array

func sortTransformedArray(nums []int, a int, b int, c int) []int {
	reverseIndex := -101
	if a != 0 {
		reverseIndex = -b / 2 * a
	}

	var reverse []int
	var normal []int
	for _, x := range nums {
		// don't reverse
		if reverseIndex == -101 {
			if b >= 0 {
				normal = append(normal, a*x*x+b*x+c)
			} else {
				reverse = append([]int{a*x*x + b*x + c}, reverse...)

			}
		} else {
			if a > 0 {
				if x >= reverseIndex {
					normal = append(normal, a*x*x+b*x+c)
				} else {
					reverse = append([]int{a*x*x + b*x + c}, reverse...)
				}
			} else {
				if x < reverseIndex {
					normal = append(normal, a*x*x+b*x+c)
				} else {
					reverse = append([]int{a*x*x + b*x + c}, reverse...)
				}
			}
		}

	}
	if a >= 0 {
		return append(reverse, normal...)
	}
	return append(normal, reverse...)
}
