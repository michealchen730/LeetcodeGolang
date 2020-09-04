package main

import (
	"math"
)

//即使照抄了代码也没看懂。。。
func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	mn, mx := nums[0], nums[0]
	for i := 0; i < len(nums); i++ {
		mn = min(mn, nums[i])
		mx = max(mx, nums[i])
	}

	//桶之间的间距
	bucketSize := max(1, (mx-mn)/(len(nums)-1))
	bucketNum := (mx-mn)/bucketSize + 1
	used := make([]bool, bucketNum)
	maxval, minval := make([]int, bucketNum), make([]int, bucketNum)
	for i := 0; i < len(maxval); i++ {
		minval[i] = math.MaxInt32
		maxval[i] = math.MinInt32
	}
	for _, v := range nums {
		bucketIdx := (v - mn) / bucketSize
		used[bucketIdx] = true
		minval[bucketIdx] = min(v, minval[bucketIdx])
		maxval[bucketIdx] = max(v, maxval[bucketIdx])
	}

	prevBucketMax, maxGap := mn, 0

	for k, _ := range used {
		if !used[k] {
			continue
		}
		//比较前一个桶的最大元素和后一个桶的最小元素的差值
		maxGap = max(maxGap, minval[k]-prevBucketMax)
		prevBucketMax = maxval[k]
	}
	return maxGap
}
