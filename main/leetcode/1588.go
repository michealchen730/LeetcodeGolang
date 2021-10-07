package main

func sumOddLengthSubarrays(arr []int) int {
	sums := make([]int, len(arr)+1)
	for k, v := range arr {
		sums[k+1] += v
	}
	res := sums[len(sums)-1]
	for t := 3; t <= len(arr); t += 2 {
		for j := t; j < len(sums); j++ {
			res += sums[j] - sums[j-t]
		}
	}
	return res
}
