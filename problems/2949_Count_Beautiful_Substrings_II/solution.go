package _949_Count_Beautiful_Substrings_II

import "math"

func eratosthenes(maxInt int) []int {
	upperBound := int(math.Sqrt(float64(maxInt)))
	isNotPrime := make([]bool, maxInt+1)
	var primes []int
	for i := 2; i <= upperBound; i++ {
		if isNotPrime[i] {
			continue
		}
		for j := i * 2; j <= maxInt; j += i {
			isNotPrime[j] = true
		}
	}
	for i := 2; i <= maxInt; i++ {
		if isNotPrime[i] {
			continue
		}
		primes = append(primes, i)
	}
	return primes
}
func beautifulSubstrings(s string, k int) int64 {
	vowels := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}
	primes := eratosthenes(k)
	// 1. 怎麼知道s[i:j]有一樣的子母音，這邊要用前綴和的概念，sum(arr[i:j]) = presum[0:j] - presum[0:i]
	// 2. 題義 (i * j) % k == 0 && i == j -> [(i - j) / 2] ** 2 = k
	// => [(i - j) / 2] = k ** (1 / 2)
	// => (i-j) = 2 * k ** (1/2)
	// 這邊有另外一個議題 如果k不為完全平方數呢？
	// 找離k最近的完全瓶芳樹 設 k = p1 * p1 * p2 則最近的完全平方數k' 為 p1 * p1 * p2 * p2
	m := 1
	for _, p := range primes {
		count := 0
		for k%p == 0 {
			count++
			k /= p
		}
		if count == 0 {
			continue
		}
		if count%2 == 1 {
			m *= int(math.Pow(float64(p), float64((count+1)/2)))
			continue
		}
		m *= int(math.Pow(float64(p), float64(count/2)))
	}
	m *= 2
	n := len(s)
	countMapping := make(map[[2]int]int, n)

	ret := int64(0)
	countMapping[[2]int{0, 0}] = 1
	t := "#" + s
	diff := 0
	for i := 1; i <= n; i++ {
		if vowels[t[i]] {
			diff++
		} else {
			diff--
		}
		ret += int64(countMapping[[2]int{diff, i % m}])
		countMapping[[2]int{diff, i % m}]++
	}
	return ret
}
