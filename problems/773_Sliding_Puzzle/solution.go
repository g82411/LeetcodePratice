package _73_Sliding_Puzzle

import "bytes"

type Queue [][]byte

func (q Queue) Len() int {
	return len(q)
}

func (q Queue) IsEmpty() bool {
	return q.Len() == 0
}

func (q *Queue) Enqueue(state []byte) {
	*q = append(*q, state)
}

func (q *Queue) Dequeue() []byte {
	prev := *q
	front := prev[0]
	*q = prev[1:]
	return front
}

func slidingPuzzle(board [][]int) int {
	// 問題在於怎麼做bfs
	targetState := "123450"
	visited := make(map[string]bool)
	var startState []byte
	const row = 2
	const col = 3
	Dirs := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			startState = append(startState, byte(board[i][j])+'0')
		}
	}
	if string(startState) == targetState {
		return 0
	}
	q := Queue{}
	q.Enqueue(startState)
	steps := 0
	for !q.IsEmpty() {
		steps++
		size := q.Len()
		for size > 0 {
			size--
			currState := q.Dequeue()
			zeroIdx := bytes.Index(currState, []byte{48})
			x := zeroIdx % col
			y := zeroIdx / col
			for _, d := range Dirs {
				dx, dy := d[0], d[1]
				tx := x + dx
				ty := y + dy
				if tx < 0 || ty < 0 {
					continue
				}
				if tx >= col || ty >= row {
					continue
				}
				nextZero := ty*col + tx
				nextState := make([]byte, 6)
				copy(nextState, currState)
				nextState[zeroIdx], nextState[nextZero] = nextState[nextZero], nextState[zeroIdx]
				key := string(nextState)

				if visited[key] {
					continue
				}
				if key == targetState {
					return steps
				}
				visited[key] = true
				q.Enqueue(nextState)
			}
		}
	}

	return -1
}
