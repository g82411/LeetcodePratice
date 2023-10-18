package _39_Word_Break

type TireNode struct {
	Next  [26]*TireNode
	IsEnd bool
}

func NewTireNode() *TireNode {
	var next [26]*TireNode
	return &TireNode{
		next,
		false,
	}
}

func wordBreak(s string, wordDict []string) bool {
	root := NewTireNode()
	// build tire
	for _, word := range wordDict {
		node := root
		for i := range word {
			char := word[i]
			if node.Next[char-'a'] == nil {
				node.Next[char-'a'] = NewTireNode()
			}
			node = node.Next[char-'a']
		}
		node.IsEnd = true
	}
	var dfs func(curr int) bool
	memoIsFail := make([]int, len(s)+5)
	dfs = func(curr int) bool { // s[curr:]
		node := root
		if curr == len(s) {
			return true
		}
		if memoIsFail[curr] == 1 {
			return false
		}
		for i := curr; i < len(s); i++ {
			if node.Next[s[i]-'a'] != nil {
				node = node.Next[s[i]-'a']
				if node.IsEnd && dfs(i+1) {
					return true
				}
			} else {
				break
			}
		}
		memoIsFail[curr] = 1
		return false
	}
	return dfs(0)
}
