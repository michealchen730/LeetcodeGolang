package main

func numPairsDivisibleBy60(time []int) int {
	sum := 0
	arr := make([]int, 60)
	for i := 0; i < len(time); i++ {
		arr[time[i]%60]++
	}
	sum += arr[0]*(arr[0]-1)/2 + arr[30]*(arr[30]-1)/2
	for i := 1; i < 30; i++ {
		sum += arr[i] * arr[60-i]
	}
	return sum
}
