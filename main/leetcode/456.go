package main

import "math"

//通过从右向左遍历的方式，找到当前数组中的最大元素
//maxK用来记录当前最大元素的右边第二大的元素
func find132pattern(nums []int) bool {
	n := len(nums)
	//candi是最后一个元素
	candidateK := []int{nums[n-1]}
	//最小值
	maxK := math.MinInt32

	//从后往前遍历
	for i := n - 2; i >= 0; i-- {
		//maxK为MinInt32的时候，很明显不会返回true
		if nums[i] < maxK {
			return true
		}

		//出栈
		for len(candidateK) > 0 && nums[i] > candidateK[len(candidateK)-1] {
			maxK = candidateK[len(candidateK)-1]
			candidateK = candidateK[:len(candidateK)-1]
		}
		if nums[i] > maxK {
			candidateK = append(candidateK, nums[i])
		}
	}
	return false

}
