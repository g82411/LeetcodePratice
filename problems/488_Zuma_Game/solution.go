package _488_Zuma_Game

import (
	"fmt"
	"math"
)

type Stat struct {
	Hand  *[5]uint8
	Board string
}

func Indexing(c byte) int {
	if c == 'R' {
		return 0
	}
	if c == 'Y' {
		return 1
	}
	if c == 'B' {
		return 2
	}
	if c == 'G' {
		return 3
	}
	return 4
}

func ReverseIndex(i int) byte {
	if i == 0 {
		return 'R'
	}
	if i == 1 {
		return 'Y'
	}
	if i == 2 {
		return 'B'
	}
	if i == 3 {
		return 'G'
	}
	return 'W'

}

func (s Stat) Encode() string {
	code := ""
	for _, v := range s.Hand {
		code += fmt.Sprint("_%d", v)
	}
	code = s.Board + "#" + code
	return code
}

func findMinStep(board string, hand string) int {
	// bruteforce
	cache := make(map[string]int)
	var eliminateContinousBoard func(board string) string
	eliminateContinousBoard = func(board string) string {
		s, e := 0, 0
		n := len(board)
		for s < n {
			for e < n && board[s] == board[e] {
				e++
			}
			if e-s >= 3 {
				next := board[:s] + board[e:]
				return eliminateContinousBoard(next)
			} else {
				s = e
			}
		}
		return board
	}
	var dfs func(stat *Stat) int
	dfs = func(stat *Stat) int {
		key := stat.Encode()
		if _, e := cache[key]; e {
			return cache[key]
		}
		ret := math.MaxInt / 2
		if len(stat.Board) == 0 {
			ret = 0
		}
		for i := range stat.Board {
			temp := stat.Board
			for j, c := range stat.Hand {
				if c == 0 {
					continue
				}
				color := ReverseIndex(j)
				needPuring := true
				if color == stat.Board[i] {
					needPuring = false
				}
				if i > 0 && stat.Board[i] == stat.Board[i-1] {
					needPuring = false
				}
				if needPuring {
					continue
				}
				nextBoard := stat.Board[:i] + string(color) + stat.Board[i:]
				stat.Board = eliminateContinousBoard(nextBoard)
				stat.Hand[j]--
				steps := dfs(stat)
				if steps != -1 {
					ret = min(ret, steps+1)
				}
				stat.Board = temp
				stat.Hand[j]++
			}
		}
		if ret == math.MaxInt/2 {
			ret = -1
		}
		cache[key] = ret
		return ret
	}
	var initHand [5]uint8
	for i := range hand {
		initHand[Indexing(hand[i])]++
	}
	initStat := &Stat{
		&initHand,
		board,
	}
	return dfs(initStat)
}
