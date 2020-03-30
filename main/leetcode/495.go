package main

func findPoisonedDuration(timeSeries []int, duration int) int {
	res := 0
	if len(timeSeries) == 0 {
		return 0
	}
	for k, _ := range timeSeries {
		if k != 0 {
			res += min(duration, timeSeries[k]-timeSeries[k-1])
		}
	}
	res += duration
	return res
}
