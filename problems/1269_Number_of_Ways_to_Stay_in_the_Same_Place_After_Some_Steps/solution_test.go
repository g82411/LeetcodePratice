package _269_Number_of_Ways_to_Stay_in_the_Same_Place_After_Some_Steps

import "testing"

// 解1 dp[k][i] = dp[k-1][i-1] + dp[k-1][i] + dp[k-1][i+1]
// 解2 -> 跑k步必須要返回，故i < K / 2 + 1 即可
func TestNumWays(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		got := numWays(3, 2)
		want := 4
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		got := numWays(2, 4)
		want := 2
		if got != want {
			t.Errorf("got %v want %v given", got, want)
		}
	})
}
