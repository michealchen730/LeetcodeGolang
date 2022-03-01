package main

import (
	"strconv"
	"strings"
)

func optimalDivision(nums []int) string {
	var div strings.Builder
	var ans strings.Builder
	ans.WriteString(strconv.Itoa(nums[0]))
	for k, v := range nums {
		if k > 0 {
			tmp := strconv.Itoa(v)
			div.WriteString(tmp)
			if k != len(nums)-1 {
				div.WriteRune('/')
			}
		}
	}
	if len(nums) > 1 {
		ans.WriteRune('/')
		t := div.String()
		if len(nums) > 2 {
			t = "(" + div.String() + ")"
		}
		ans.WriteString(t)
	}
	return ans.String()
}
