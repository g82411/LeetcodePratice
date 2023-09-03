package _89_Robot_Room_Cleaner

import "fmt"

type Robot struct {
}

func (robot *Robot) Move() bool {
	return false
}

func (robot *Robot) TurnLeft()  {}
func (robot *Robot) TurnRight() {}

func (robot *Robot) Clean() {}

// 太鬼了, 這題為什麼沒有測試呢 因為他給你API要你盲掃阿
// 不過想想也是 很多開發中基本上只提供核心API 但不會允許去改

func cleanRoom(robot *Robot) {
	DIRECTIONS := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	visited := make(map[string]bool)
	encodeCoordinate := func(x, y int) string {
		s := fmt.Sprintf("%v_%v", x, y)
		return s
	}
	turnAround := func() {
		robot.TurnLeft()
		robot.TurnLeft()
		robot.Move()
		robot.TurnLeft()
		robot.TurnLeft()
	}
	var dfs func(x, y, direction int)
	dfs = func(x, y, direction int) {
		robot.Clean()
		for k := 1; k <= 4; k++ {
			robot.TurnRight()
			nextDirectionIdx := (direction + k) % 4
			nextX := x + DIRECTIONS[nextDirectionIdx][0]
			nextY := y + DIRECTIONS[nextDirectionIdx][1]
			code := encodeCoordinate(nextX, nextY)
			if !visited[code] && robot.Move() {
				visited[code] = true
				dfs(nextX, nextY, nextDirectionIdx)
				turnAround()
			}
		}
	}
	visited[encodeCoordinate(0, 0)] = true
	dfs(0, 0, 0)
}
