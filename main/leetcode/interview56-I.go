package main

func singleNumbers(nums []int) []int {
	//拿到异或和
	sum := 0
	for _, v := range nums {
		sum ^= v
	}
	//从最低位开始，找到位置为0的第一个数
	//正确性：1&3=1 二进制：1（1）， 11（3）;2&6=2 二进制：10（2）， 110（6）
	t := 1
	for t&sum != t {
		t <<= 1
	}
	a, b := 0, 0
	for _, v := range nums {
		if v&t == 0 {
			a ^= v
		} else {
			b ^= v
		}
	}
	return []int{a, b}
}
