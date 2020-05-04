package main

//更新版本的解法
func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	//设置起点和终点（注意这里终点是达不到的）
	start, end := 0, 1
	//设置此时的步数，此时至少需要一步才能走完
	steps := 1
	//mx表示在steps为具体某个i时能走的最大步数，会在循环内更新
	mx := 0
	for start < end {
		temp := start + nums[start]
		if temp >= len(nums)-1 { //表示当前start位置能达到终点
			return steps
		}
		if temp > mx { //这里更新mx
			mx = temp
		}
		if start == end-1 { //当start走到end-1时，需要更新end，end就是mx+1的值
			end = mx + 1
			steps++ //需要更新步数，因为当前steps已经走完了
		}
		start++
	}
	return 0 //这里的return其实是不会返回的
}

//最初版本的解法
func jump45(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	arr := make([]int, len(nums))
	arr[0] = 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j <= i+nums[i]; j++ {
			if j < len(arr) && (arr[j] == 0 || min(arr[j], arr[i]+1) == arr[i]+1) {
				arr[j] = arr[i] + 1
			}
		}
	}
	return arr[len(arr)-1]
}
func min(x int, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}
