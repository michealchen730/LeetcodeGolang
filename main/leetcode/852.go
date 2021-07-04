package main

//func peakIndexInMountainArray(arr []int) int {
//	for i:=1;i<len(arr)-1;i++{
//		if arr[i]>arr[i-1]&&arr[i]>arr[i+1]{
//			return i
//		}
//	}
//	return 0
//}

func peakIndexInMountainArray(arr []int) int {
	low, high := 1, len(arr)-2
	for low < high {
		mid := low + (high-low)/2
		if arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
			return mid
		} else if arr[mid] > arr[mid-1] && arr[mid] < arr[mid+1] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}
