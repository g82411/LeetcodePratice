package _222_Number_of_Ways_to_Select_Buildings

func numberOfWays(s string) int64 {
	// dp -> 前n種狀態
	// 出現一個0
	// 有可能10 -> 10 = 1出現的次數
	// 010 -> 01 -> 0
	// 101 -> 10 -> 1
	occur01 := 0
	occur10 := 0
	occur1 := 0
	occur0 := 0
	occur101 := 0
	occur010 := 0

	for i := range s {
		if s[i] == '0' {
			occur0++
			occur10 += occur1
			occur010 += occur01
			continue
		}
		occur1++
		occur01 += occur0
		occur101 += occur10
	}
	return int64(occur101 + occur010)
}
