package main

//这题我不会想到枚举k（自己会觉得时间复杂度过高）
func maxNumber(nums1 []int, nums2 []int, k int) []int {
	res := make([]int, k)

	for i := 0; i <= k; i++ {
		arr1 := MonotonicStack(nums1, i)
		arr2 := MonotonicStack(nums2, k-i)
		if len(arr1)+len(arr2) == k {
			arr3 := mergeArr(arr1, arr2)
			res = maxArr(arr3, res)
		}
	}
	return res
}

func maxArr(arr1, arr2 []int) []int {
	for k, v := range arr1 {
		if v > arr2[k] {
			return arr1
		} else if v < arr2[k] {
			return arr2
		}
	}
	return arr1
}

func MonotonicStack(arr []int, k int) []int {
	if k == 0 || k > len(arr) {
		return nil
	}

	var stack []int

	for i := 0; i < len(arr); i++ {
		if len(stack) == 0 || len(arr)-i+len(stack) == k {
			stack = append(stack, arr[i])
			continue
		}
		for len(stack) > 0 && len(stack)+len(arr)-i > k && stack[len(stack)-1] < arr[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) < k {
			stack = append(stack, arr[i])
		}
	}
	return stack
}

func mergeArr(nums1 []int, nums2 []int) []int {
	res := make([]int, 0, len(nums1)+len(nums2))
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] > nums2[j] {
			res = append(res, nums1[i])
			i++
		} else if nums1[i] < nums2[j] {
			res = append(res, nums2[j])
			j++
		} else {
			res = append(res, nums1[i])
			//例子
			//[6,7]
			//[6,0,7]
			m, n := i, j
			for m < len(nums1) && n < len(nums2) {
				if nums1[m] > nums2[n] {
					i++
					break
				} else if nums1[m] < nums2[n] {
					j++
					break
				} else {
					m++
					n++
				}
			}
			//例子
			//[6,7]
			//[6,7]
			if m == len(nums1) && n == len(nums2) {
				i++
				continue
			}
			//[6]
			//[6,7]
			//让长的一方通行
			if n == len(nums2) {
				i++
				continue
			}
			if m == len(nums1) {
				j++
			}
		}
	}
	if i == len(nums1) {
		res = append(res, nums2[j:]...)
	} else {
		res = append(res, nums1[i:]...)
	}
	return res
}
