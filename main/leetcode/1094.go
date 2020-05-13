package main

func carPooling(trips [][]int, capacity int) bool {
	stations := make([]int, 1001)
	for _, v := range trips {
		for i := v[1]; i < v[2]; i++ {
			stations[i] += v[0]
			if stations[i] > capacity {
				return false
			}
		}
	}
	return true
}
