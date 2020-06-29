package main

//堆排和快排
//这里快排的思路真的很厉害
func findKthLargest(nums []int, k int) int {
	var arr []int
	for i := 0; i < k; i++ {
		arr = append(arr, nums[i])
	}
	buildMinHeap215(arr)
	for i := k; i < len(nums); i++ {
		if nums[i] > arr[0] {
			arr[0] = nums[i]
			adjustMinHeap215(arr, 0, k)
		}
	}
	return arr[0]
}

func buildMinHeap215(arr []int) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		adjustMinHeap215(arr, i, len(arr))
	}
}
func adjustMinHeap215(arr []int, i int, length int) {
	for k := 2*i + 1; k < length; k = 2*k + 1 {
		if k+1 < length && (arr[k] > arr[k+1]) {
			k = k + 1
		}
		if arr[k] < arr[i] {
			arr[k], arr[i] = arr[i], arr[k]
			i = k
		} else {
			break
		}
	}
}
