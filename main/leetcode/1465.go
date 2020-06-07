package main

import "sort"

func maxArea1465(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)
	mr, mc := horizontalCuts[0]-0, verticalCuts[0]-0
	for i := 1; i < len(horizontalCuts); i++ {
		mr = max(mr, horizontalCuts[i]-horizontalCuts[i-1])
	}
	for j := 1; j < len(verticalCuts); j++ {
		mc = max(mc, verticalCuts[j]-verticalCuts[j-1])
	}
	mr = max(mr, h-horizontalCuts[len(horizontalCuts)-1])
	mc = max(mc, w-verticalCuts[len(verticalCuts)-1])
	return (mr % 1000000007) * (mc % 1000000007) % 1000000007
}
