package main

//枚举思路
func findKthPositive(arr []int, k int) int {
	mp := make(map[int]bool)
	for _, v := range arr {
		mp[v] = true
	}
	t := 1
	cnt := 0
	for true {
		if !mp[t] {
			cnt++
			if cnt == k {
				return t
			}
		}
		t++
	}
	return -1
}

//二分（推荐思路）
