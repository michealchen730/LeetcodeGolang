package main

func minSumOfLengths(arr []int, target int) int {
	arr2 := make([]int, len(arr)+1)
	mp := make(map[int]int)
	res := []int{0, 0}
	str := []int{0, 0}
	mp[0] = 0
	for i := 1; i < len(arr2); i++ {
		arr2[i] += arr2[i-1] + arr[i-1]
		mp[arr2[i]] = i
		if v, ok := mp[arr2[i]-target]; ok {
			//v实际上是起点
			if res[0] == 0 {
				str[0] = v
				res[0] = i - v
				continue
			}
			//判断第二个
			if res[1] == 0 {
				if v+1 > str[0]+res[0] {
					res[1] = i - v
					str[1] = v
				} else if i-v < res[0] {
					res[0] = i - v
					str[1] = v
				}
				continue
			}
			if res[0] == 1 && res[1] == 1 {
				return 2
			}
			//手动排序
			if res[0] > res[1] {
				res[0], res[1] = res[1], res[0]
				str[0], str[1] = str[1], str[0]
			}
			if i-v < res[1] {
				if v+1 > res[0]+str[0] {
					res[1] = i - v
					str[1] = v
				} else if i-v < res[0] {
					if v+1 > res[1]+str[1] {
						res[0] = i - v
						str[0] = v
					}
				}
			}
		}
	}
	if res[0] == 0 || res[1] == 0 {
		return -1
	}
	return res[0] + res[1]
}
