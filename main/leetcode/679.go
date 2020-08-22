package main

import (
	"math"
)

//虽然是1到9数字的牌
//但在一轮计算后，还是会出现0这一数的
func judgePoint24(nums []int) bool {
	deviation := 1e-6
	solved := false
	var dfs func(nums []float64)
	dfs = func(nums []float64) {
		if solved {
			return
		}
		if len(nums) == 1 {
			if math.Abs(nums[0]-24) < deviation {
				solved = true
			}
			return
		}
		for i := 0; i < len(nums)-1; i++ {
			for j := i + 1; j < len(nums); j++ {
				//乘法和加法具有交换律
				var rea [6]float64
				rea[0], rea[1], rea[2], rea[3] = nums[i]*nums[j], nums[i]+nums[j], nums[i]-nums[j], nums[j]-nums[i]
				//特判除法
				if nums[i] != 0.0 {
					rea[4] = nums[j] / nums[i]
				}
				if nums[j] != 0.0 {
					rea[5] = nums[i] / nums[j]
				}
				newArr := []float64{0.0}
				for e := 0; e < len(nums); e++ {
					if e != i && e != j {
						newArr = append(newArr, nums[e])
					}
				}
				for k := 0; k < len(rea); k++ {
					newArr[0] = rea[k]
					dfs(newArr)
				}
			}
		}
	}
	var arr []float64
	for i := 0; i < 4; i++ {
		arr = append(arr, float64(nums[i]))
	}
	dfs(arr)
	return solved
}
