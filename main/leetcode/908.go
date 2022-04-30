package main

//找到数组中的最大值和最小值
//判断这两个值的差值是否小于2*k
func smallestRangeI(nums []int, k int) int {
	minNum, maxNum := 10000, 0
	for _, v := range nums {
		if v < minNum {
			minNum = v
		}
		if v > maxNum {
			maxNum = v
		}
	}
	if maxNum-minNum <= 2*k {
		return 0
	} else {
		return maxNum - minNum - 2*k
	}
}
