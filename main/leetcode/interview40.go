package main

func getLeastNumbers(arr []int, k int) []int {
	if k != 0 {
		buildMaxHeap(arr[0:k])
		for i := k; i < len(arr); i++ {
			if arr[i] < arr[0] {
				arr[0] = arr[i]
				adjustHeap(arr, 0, k)
			}
		}
	}
	return arr[0:k]
}

func heapSort(arr []int) {
	buildMaxHeap(arr)
	for i := len(arr) - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		adjustHeap(arr, 0, i)
	}
}

//最大堆的建立
func buildMaxHeap(arr []int) {
	//因为堆是完全二叉树
	//所以只要对每个非叶子结点全部调整一遍即可
	for i := len(arr)/2 - 1; i >= 0; i-- {
		adjustHeap(arr, i, len(arr))
	}
}
func adjustHeap(arr []int, i int, length int) {
	for k := 2*i + 1; k < length; k = 2*k + 1 {
		if k+1 < length && arr[k] < arr[k+1] {
			k = k + 1
		}
		if arr[k] > arr[i] {
			arr[k], arr[i] = arr[i], arr[k]
			i = k
		} else {
			break
		}
	}
}
