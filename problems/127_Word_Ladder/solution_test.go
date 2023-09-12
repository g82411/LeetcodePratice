package _27_Word_Ladder

import "testing"

func TestLadderLength(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		beginWord := "hit"
		endWord := "cog"
		wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
		expect := 5
		actual := ladderLength(beginWord, endWord, wordList)
		if expect != actual {
			t.Errorf("Expect: %v, actual: %v", expect, actual)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		beginWord := "hit"
		endWord := "cog"
		wordList := []string{"hot", "dot", "dog", "lot", "log"}
		expect := 0
		actual := ladderLength(beginWord, endWord, wordList)
		if expect != actual {
			t.Errorf("Expect: %v, actual: %v", expect, actual)
		}
	})

}
