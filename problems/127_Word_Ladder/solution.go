package _27_Word_Ladder

type Queue []string

func (q Queue) Len() int {
	return len(q)
}

func (q Queue) IsEmpty() bool {
	return q.Len() == 0
}

func (q *Queue) Push(s string) {
	*q = append(*q, s)
}

func (q *Queue) Pop() string {
	prev := *q
	first := prev[0]
	*q = prev[1:]
	return first
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	POOLS := []byte{
		'a', 'b', 'c', 'd', 'e', 'f', 'g',
		'h', 'i', 'j', 'k', 'l', 'm', 'n',
		'o', 'p', 'q', 'r', 's', 't', 'u',
		'v', 'w', 'x', 'y', 'z',
	}
	candidate := make(map[string]bool)
	for _, s := range wordList {
		candidate[s] = true
	}
	q := &Queue{}
	if _, e := candidate[endWord]; !e {
		return 0
	}
	steps := 0
	q.Push(beginWord)
	for !q.IsEmpty() {
		steps++
		for s := q.Len(); s > 0; s-- {
			curWord := q.Pop()
			for p := 0; p < len(curWord); p++ {
				for c := 0; c < len(POOLS); c++ {
					if curWord[p] == POOLS[c] {
						continue
					}
					w := curWord[:p] + string(POOLS[c]) + curWord[p+1:]
					if w == endWord {
						return steps + 1
					}
					if _, e := candidate[w]; !e {
						continue
					}
					delete(candidate, w)
					q.Push(w)
				}
			}
		}

	}
	return 0
}
