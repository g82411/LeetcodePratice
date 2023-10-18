package _40_Word_Break_II

import "strings"

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

func wordBreak(s string, wordDict []string) []string {
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
	var dfs func(curr int, words []string) bool
	memoIsFail := make([]int, len(s)+5)
	var words []string
	var res []string
	dfs = func(curr int, words []string) bool { // s[curr:]
		node := root
		if curr == len(s) {
			t := strings.Join(words, " ")
			res = append(res, t)
			return true
		}
		if memoIsFail[curr] == 1 {
			return false
		}
		flag := false
		for i := curr; i < len(s); i++ {
			if node.Next[s[i]-'a'] != nil {
				node = node.Next[s[i]-'a']
				if node.IsEnd {
					words = append(words, s[curr:i+1])
					if dfs(i+1, words) {
						flag = true
					}
					words = words[:len(words)-1]
				}
			} else {
				break
			}
		}
		if !flag {
			memoIsFail[curr] = 1
		}
		return flag
	}
	dfs(0, words)
	return res
}
