package main

//暴力的基础上用数学方法进行了一轮优化就可以得到答案
func totalHammingDistance(nums []int) int {
	res := 0
	//枚举30位
	for i := 0; i < 30; i++ {
		cnt := 0
		for _, v := range nums {
			cnt += v >> i & 1
		}
		res += cnt * (len(nums) - cnt)
	}
	return res
}
