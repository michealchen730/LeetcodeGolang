package main

func minArray(numbers []int) int {
	low, high := 0, len(numbers)-1
	for low < high {
		mid := low + (high-low)/2
		if numbers[mid] == numbers[high] {
			high--
			continue
		}
		if numbers[mid] < numbers[high] {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return numbers[low]
}
