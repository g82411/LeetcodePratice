package _862_Maximum_Element_Sum_of_a_Complete_Subset_of_Indices

//先推倒 sqrt (a * b) belong nature 若且維若 a * b = a * (i * i) * a * (j * j)
// 固定a = 1 ~ n
// 則 a= n 時sum = nums[n * i * i] + nums[n * j * j] + ...
// 又 max(n) = 100 故 i 最高到9

func maximumSum(nums []int) int64 {
	ret := int64(0)
	n := len(nums)
	for k := 1; k <= n; k++ {
		sum := int64(0)
		for i := 1; k*i*i <= n; i++ {
			sum += int64(nums[k*i*i-1])
		}
		ret = max(ret, sum)
	}
	return ret
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
