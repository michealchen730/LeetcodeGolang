package main

func findSolution(customFunction func(int, int) int, z int) [][]int {
	var res [][]int
	for i := 1; i <= 1000; i++ {
		low, high := 1, 1000
		t1, t2 := customFunction(i, low), customFunction(i, high)
		if z < t1 || z > t2 {
			continue
		}
		for low <= high {
			mid := low + (high-low)/2
			t3 := customFunction(i, mid)
			if z == t3 {
				res = append(res, []int{i, mid})
				break
			} else if z > t3 {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return res
}
