package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	var res []string
	if len(nums) == 0 {
		return res
	}

	start := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 {
			if i == len(nums)-1 {
				res = append(res, strconv.Itoa(start)+"->"+strconv.Itoa(nums[i]))
				return res
			}
			continue
		} else {
			if nums[i-1] == start {
				res = append(res, strconv.Itoa(start))
			} else {
				res = append(res, strconv.Itoa(start)+"->"+strconv.Itoa(nums[i-1]))
			}
			start = nums[i]
		}
	}
	res = append(res, strconv.Itoa(nums[len(nums)-1]))
	return res
}

func main() {
	fmt.Println(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
}
