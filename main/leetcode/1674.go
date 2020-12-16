package main

//我自己的想法，砸了
//最后还是使用题解区大佬的代码
func minMoves1674(nums []int, limit int) int {
	diff := make([]int, 2*limit+2)

	n := len(nums)
	for i := 0; i < n/2; i++ {
		a, b := nums[i], nums[n-1-i]

		l, r := 2, 2*limit
		diff[l] += 2
		diff[r+1] -= 2

		l = 1 + min(a, b)
		r = limit + max(a, b)
		diff[l] += -1
		diff[r+1] -= -1

		l = a + b
		r = a + b
		diff[l] += -1
		diff[r+1] -= -1
	}

	res := n
	sum := 0
	for i := 2; i <= 2*limit; i++ {
		sum += diff[i]
		if sum < res {
			res = sum
		}
	}
	return res
}
