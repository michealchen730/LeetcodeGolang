package main

import (
	"sort"
)

//func main()  {
//	fmt.Println(threeSumClosest([]int{0,2,1,-3},1))
//}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	max := nums[0] + nums[1] + nums[2] - target
	for i := 0; i < len(nums)-2; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, len(nums)-1
		for j < k {
			temp := nums[i] + nums[j] + nums[k] - target
			if abs(temp) < abs(max) {
				max = temp
			}
			if temp > 0 {
				k--
			}
			if temp < 0 {
				j++
			}
			if temp == 0 {
				return target
			}
		}
	}
	return target + max
}
