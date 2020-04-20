package main

import "fmt"

func numberOfSubarrays(nums []int, k int) int {
	var stack []int
	for k, v := range nums {
		if v%2 != 0 {
			stack = append(stack, k)
		}
	}
	res := 0
	if len(stack) >= k {
		i, j, left, right := 0, k-1, 0, 0
		for j < len(stack) {
			if i == 0 {
				left = stack[i] - 0 + 1
			} else {
				left = stack[i] - stack[i-1]
			}
			if j != len(stack)-1 {
				right = stack[j+1] - stack[j]
			} else {
				right = len(nums) - stack[j]
			}
			res, i, j = res+left*right, i+1, j+1
		}
	}
	return res
}

func main() {
	fmt.Println(numberOfSubarrays([]int{2, 2, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 3))
}
