package main

func createTargetArray(nums []int, index []int) []int {
	res := []int{nums[0]}
	for i := 1; i < len(index); i++ {
		if index[i] == len(res) {
			res = append(res, nums[i])
		} else {
			res = append(res[0:index[i]], append([]int{nums[i]}, res[index[i]:]...)...)
		}
	}
	return res
}
