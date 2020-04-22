package main

func minNumberOfFrogs(croakOfFrogs string) int {
	arr := make([]int, 5)
	mp := map[int32]int{'c': 0, 'r': 1, 'o': 2, 'a': 3, 'k': 4}
	mx := 0
	for _, v := range croakOfFrogs {
		arr[mp[v]]++
		for i := 1; i < len(arr); i++ {
			if arr[i] > arr[i-1] {
				return -1
			}
		}
		if v == 'k' {
			if arr[0] > mx {
				mx = arr[0]
			}
			for i := 0; i < len(arr); i++ {
				arr[i]--
			}
		}
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			return -1
		}
	}
	return mx
}
