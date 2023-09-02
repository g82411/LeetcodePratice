package _419_Minimum_Number_of_Frogs_Croaking

// 題目問的是Croak算做一次青蛙叫
// 如果今天出現CCRROOAAKK算是幾次青蛙叫
// 並且是求最少 所以CROAKCROAK = 1
func minNumberOfFrogs(croakOfFrogs string) int {
	// 看到CC就至少要兩隻, 再來確定他有沒有叫完
	// 如果CRC就算是基因突變 要return -1
	c := 0
	r := 0
	o := 0
	a := 0
	ret := 0
	for i := 0; i < len(croakOfFrogs); i++ {
		if croakOfFrogs[i] == 'c' {
			c++
		} else if croakOfFrogs[i] == 'r' {
			c--
			r++
			if c < 0 {
				return -1
			}
		} else if croakOfFrogs[i] == 'o' {
			r--
			o++
			if r < 0 {
				return -1
			}
		} else if croakOfFrogs[i] == 'a' {
			o--
			a++
			if o < 0 {
				return -1
			}
		} else if croakOfFrogs[i] == 'k' {
			a--
			if a < 0 {
				return -1
			}
		}
		t := c + r + o + a
		if t > ret {
			ret = t
		}
	}
	if c+r+o+a > 0 {
		return -1
	}
	return ret
}
