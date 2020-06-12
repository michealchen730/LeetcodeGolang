package main

//此题可以二重循环、哈希、建立一个数组以及位运算，这里仅给出位运算

//位运算解法
func isUnique(astr string) bool {
	fr, sc := int64(0), int64(0)
	for _, v := range astr {
		if v >= 64 {
			t := int64(1) << (v - 64)
			if t&sc != 0 {
				return false
			} else {
				sc |= t
			}
		} else {
			t := int64(1) << v
			if t&fr != 0 {
				return false
			} else {
				fr |= t
			}
		}
	}
	return true
}
