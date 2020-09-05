package main

func threeConsecutiveOdds(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] % 2
		if i > 2 && arr[i] == 1 && arr[i-1] == 1 && arr[i-2] == 1 {
			return true
		}
	}
	return false
}
