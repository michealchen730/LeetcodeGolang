package main

//left,right数组记录每个i点为山顶分别可以向左、向右走多远
func longestMountain(A []int) int {
	left, right := make([]int, len(A)), make([]int, len(A))
	for i := 1; i < len(A); i++ {
		if A[i] > A[i-1] {
			left[i] = left[i-1] + 1
		}
	}
	for i := len(A) - 2; i >= 0; i-- {
		if A[i] > A[i+1] {
			right[i] = right[i+1] + 1
		}
	}
	res := 0
	for i := 1; i < len(A); i++ {
		if left[i] == 0 || right[i] == 0 {
			continue
		}
		mnt := left[i] + right[i] + 1
		if mnt > 2 {
			res = max(res, mnt)
		}
	}
	return res
}

//right数组可以省略，节省空间复杂度
func longestMountain2(A []int) int {
	left := make([]int, len(A))
	for i := 1; i < len(A); i++ {
		if A[i] > A[i-1] {
			left[i] = left[i-1] + 1
		}
	}
	res := 0
	right := 0
	for i := len(A) - 2; i >= 0; i-- {
		if A[i] > A[i+1] {
			right++
			if left[i] != 0 {
				res = max(res, left[i]+right+1)
			}
		} else {
			right = 0
		}

	}
	return res
}
