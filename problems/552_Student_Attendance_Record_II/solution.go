package _52_Student_Attendance_Record_II

func checkRecord(n int) int {
	// 難就難在狀態轉移
	// 缺席有兩種 0 -> 1
	// 遲到有3種 0 -> 1 -> 2
	// 所以 狀態有6種 00 01 02 10 11 12
	// then x0 -> x1 -> x2 0y -> 1y
	mod := int(1e9) + 7
	d00 := 1
	d01 := 0
	d02 := 0
	d10 := 0
	d11 := 0
	d12 := 0
	for i := 0; i < n; i++ {
		td01 := d01
		td02 := d02
		td10 := d10
		td11 := d11
		td12 := d12
		td00 := d00
		// 前兩天遲到今天也可以不遲到
		d00 = (td00 + td02 + td01) % mod
		// 連續遲到只能發生在前一天遲到的狀況
		d01 = td00
		d02 = td01
		// 昨天遲到 今天缺席 + 昨天缺席 今天沒有缺席
		d10 = (td00 + td01 + td02 + td10 + td11 + td12) % mod
		d11 = td10
		d12 = td11
	}
	return (d00 + d01 + d02 + d11 + d12 + d10) % mod
}
