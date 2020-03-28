package main

func hIndex(citations []int) int {
	low, high, res := 0, len(citations)-1, 0
	for low <= high {
		mid := low + (high-low)/2
		if citations[len(citations)-mid-1] >= mid+1 {
			if mid+1 > res {
				res = mid + 1
			}
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return res
}
