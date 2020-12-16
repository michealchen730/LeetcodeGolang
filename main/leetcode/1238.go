package main

//格雷编码，这里是它的变种
//格雷编码是一个环，第一个和第n个也只差一个位的不同
func circularPermutation(n int, start int) []int {
	t := 1
	res := make([]int, 1<<n)
	for n >= 1 {
		temp := t
		for k := t; k < 2*t; k++ {
			res[k] = res[temp-1] ^ t
			temp--
		}
		t = t << 1
		n--
	}

	for k, v := range res {
		if v == start {
			res = append(res, res[:k]...)
			res = res[k:]
			break
		}
	}
	return res
}
