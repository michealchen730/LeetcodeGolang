package main

//这题的测试样例会导致常规思路超时
func canConvertString(s string, t string, k int) bool {
	if len(s) != len(t) {
		return false
	}

	//不管一个字母被替换了多少次，都可以%26
	//比如一个字母被替换了27次，那就是等同于被替换1次

	//举例：当k=27次，那么“等同替换”中，1次和27次都等同于1次，其他2-26都只能使用1次
	avgCnt := k / 26
	mxSize := k % 26
	arr := make([]int, 27)

	for i := 1; i < 27; i++ {
		arr[i] = avgCnt
		if i <= mxSize {
			arr[i] += 1
		}
	}

	cnt := make([]int, 27)

	for i := 0; i < len(s); i++ {
		temp := int(t[i]) - int(s[i])
		if temp != 0 {
			if temp < 0 {
				temp += 26
			}
			//最后的测试样例就是要我们判断，一旦某个字母等同替换的次数超出允许了，就可以直接返回false
			cnt[temp]++
			if cnt[temp] > arr[temp] {
				return false
			}
		}
	}
	return true
}
