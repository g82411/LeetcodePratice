package _785_Sort_Vowels_in_a_String

// 就簡單的用pq解
// 1 pass把所有的vowel放進pq
// 2 pass把所有的vowel pop出來

import "container/heap"

type PQ []byte

func (q PQ) Len() int {
	return len(q)
}

func (q PQ) IsEmpty() bool {
	return q.Len() == 0
}

func (q PQ) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q PQ) Less(i, j int) bool {
	return q[i] < q[j]
}

func (q *PQ) Push(e interface{}) {
	*q = append(*q, e.(byte))
}

func (q *PQ) Pop() interface{} {
	prev := *q
	n := len(prev)
	top := prev[n-1]
	*q = prev[0 : n-1]
	return top
}

func sortVowels(S string) string {
	dic := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}
	q := PQ{}
	heap.Init(&q)

	isVowel := func(c byte) bool {
		return dic[c]
	}
	for i := range S {
		if isVowel(S[i]) {
			heap.Push(&q, S[i])
		}
	}
	var ans []byte
	for i := range S {
		if isVowel(S[i]) {
			ans = append(ans, heap.Pop(&q).(byte))
			continue
		}
		ans = append(ans, S[i])
	}
	return string(ans)
}
