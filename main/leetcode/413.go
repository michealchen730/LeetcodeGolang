package main

func numberOfArithmeticSlices(nums []int) int {
	//表示当前等差数列的长度
	tmp := 0
	res := 0
	for k, v := range nums {
		//判断递增
		if tmp < 2 || (k > 1 && v-nums[k-1] == nums[k-1]-nums[k-2]) {
			tmp++
			continue
		}
		//此时一定不递增
		if tmp >= 3 {
			res += (tmp - 1) * (tmp - 2) / 2
		}
		tmp = 2
	}
	if tmp >= 3 {
		res += (tmp - 1) * (tmp - 2) / 2
	}
	return res
}
