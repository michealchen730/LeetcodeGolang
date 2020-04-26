package main

func reversePairs(nums []int) int {
	return mergeI51(nums, 0, len(nums)-1)
}

func mergeI51(arr []int, s, e int) int {
	if s >= e {
		return 0
	}
	mid := s + (e-s)/2
	//统计左右逆序对的数目
	res := mergeI51(arr, s, mid) + mergeI51(arr, mid+1, e)
	//统计左右合并时逆序对的个数
	var temp []int
	i, j := s, mid+1
	for i <= mid && j <= e {
		if arr[i] <= arr[j] {
			temp = append(temp, arr[i])
			res += j - (mid + 1)
			i++
		} else {
			temp = append(temp, arr[j])
			j++
		}
	}
	//如果左边数组有剩余
	for ; i <= mid; i++ {
		temp = append(temp, arr[i])
		res += e - mid
	}
	//如果右边数组有剩余
	for ; j <= e; j++ {
		temp = append(temp, arr[j])
	}
	for k := s; k <= e; k++ {
		arr[k] = temp[k-s]
	}
	return res
}
