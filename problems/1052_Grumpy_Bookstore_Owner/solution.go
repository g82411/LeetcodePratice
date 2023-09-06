package _052_Grumpy_Bookstore_Owner

func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	mSum := 0
	n := len(customers)
	for i := 0; i < n; i++ {
		if grumpy[i] == 0 {
			mSum += customers[i]
		}
	}
	ret := 0
	for i := 0; i < n; i++ {
		if grumpy[i] == 1 {
			mSum += customers[i]
		}
		if i >= minutes && grumpy[i-minutes] == 1 {
			mSum -= customers[i-minutes]
		}
		ret = max(ret, mSum)
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
