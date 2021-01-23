package main

func lastStoneWeight(stones []int) int {
	arr := make([]int, 1001)
	mx := 0
	for _, v := range stones {
		arr[v]++
		mx = max(mx, v)
	}
	for true {
		w1, w2 := 0, 0
		for k := mx; k >= 0; k-- {
			if arr[k] > 0 {
				w1 = k
				mx = k
				arr[k] -= 1
				break
			}
		}
		for k := mx; k >= 0; k-- {
			if arr[k] > 0 {
				w2 = k
				arr[k] -= 1
				break
			}
		}
		if w1 == 0 {
			return 0
		}
		if w2 == 0 {
			return w1
		} else {
			arr[w1-w2]++
		}
	}
	return -1
}
