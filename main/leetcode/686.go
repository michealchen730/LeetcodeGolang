package main

func repeatedStringMatch(a string, b string) int {
	str := a
	num := 0
	//拿到所有num值
	for i, _ := range b {
		num += int(b[i] - 'a')
	}
	for len(str) < len(b)+2*len(a) {
		if subString(str, b, num) {
			return len(str) / len(a)
		}
		str += a
	}
	return -1
}

//这里的rabin-karp算法对于哈希值的计算只是重复的累加
func subString(long, short string, num int) bool {
	if len(long) < len(short) {
		return false
	}
	n := len(short)
	temp := 0
	for i := 0; i < n; i++ {
		temp += int(long[i] - 'a')
	}
	//rabin-karp哈希值相同并不表示存在子串，依旧需要进行验证
	if temp == num && long[:n] == short {
		return true
	}

	for i := 0; i+n < len(long); i++ {
		j := i + n
		temp -= int(long[i] - 'a')
		temp += int(long[j] - 'a')
		if temp == num && long[i+1:j+1] == short {
			return true
		}
	}
	return false
}
