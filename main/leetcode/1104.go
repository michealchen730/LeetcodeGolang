package main

func pathInZigZagTree(label int) []int {
	nums := 0
	arr := []int{1}
	sum := 1
	for arr[len(arr)-1] < label {
		t := arr[len(arr)-1]
		sum *= 2
		arr = append(arr, t+sum)
		nums++
	}
	res := []int{label}
	k := len(arr) - 1
	for res[0] != 1 {
		t, tar := res[0], 0
		if nums%2 == 0 {
			m := (t-arr[k-1]-1)/2 + 1
			tar = arr[k-1] - m + 1
		} else {
			m := (arr[k] - t + 2) / 2
			tar = (arr[k-1]+1)/2 + m - 1
		}
		res = append([]int{tar}, res...)
		k--
		nums--
	}
	return res
}
