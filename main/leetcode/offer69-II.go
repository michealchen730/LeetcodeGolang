package main

func peakIndexInMountainArray69(arr []int) int {
	l, r := 1, len(arr)-2
	for l < r {
		mid := l + (r-l)/2
		if arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
			return mid
		}
		if arr[mid] < arr[mid+1] {
			l = mid + 1
		}
		if arr[mid-1] > arr[mid] {
			r = mid - 1
		}
	}
	return l
}
