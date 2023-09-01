package _094_Car_Pooling

func carPooling(trips [][]int, capacity int) bool {
	diff := make([]int, 1001)
	for _, trip := range trips {
		passengers, start, end := trip[0], trip[1], trip[2]
		diff[start] += passengers
		diff[end] -= passengers
	}
	totalPassengersInSameTime := 0
	for i := 0; i < len(diff); i++ {
		totalPassengersInSameTime += diff[i]
		if totalPassengersInSameTime > capacity {
			return false
		}
	}
	return true
}
