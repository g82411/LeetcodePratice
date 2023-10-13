package _61_Bomb_Enemy

import "testing"

// 同一列 如果沒卡牆 那be 都會一樣
// 如果卡牆 那就從牆後開始算

func TestMaxKilledEnemies(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		expected := 3
		actual := maxKilledEnemies([][]byte{
			{'0', 'E', '0', '0'},
			{'E', '0', 'W', 'E'},
			{'0', 'E', '0', '0'},
		})
		if actual != expected {
			t.Errorf("got %v want %v given", actual, expected)
		}
	})
}
