package main

import "math"

//大佬的思路简洁牛批
//func sumSubarrayMins(arr []int) int {
//	//单调栈
//	stack:=[]int{-1}
//	sum:=0
//	stacknum:=0
//	//以f(idx)表示以第idx个元素为右边界连续子数组的和，我们需要求f(4)
//	//以4为右边界的连续子数组总共有4个
//	//先假设4前面的元素是2,此时数组为[3,1,2,4]
//	//此时f(3)表示以2为右边界的连续子数组的最小值之和
//	//4>2，那么除了[4]以外的[2,4],[1,2,4]以及[3,1,2,4]三个连续子数组的和就是f(3)
//	//假设4前面的元素是5，此时数组为[3,1,5,4]
//	//4<5，那么除了[4],[4,5]以外的[1,5,4],[3,1,5,4]两个连续子数组的和就是f(2)
//	div:=int(math.Pow(10,9)+7)
//	for k,v:=range arr{
//		//单调栈进栈
//		for len(stack)>1&&v<arr[stack[len(stack)-1]]{
//			//pop过程
//			pop:=stack[len(stack)-1]
//			stack = stack[:len(stack)-1]
//			//
//			stacknum-=arr[pop]*(pop-stack[len(stack)-1])
//		}
//		stack = append(stack,k)
//		stacknum+=v*(k-stack[len(stack)-2])
//		sum+=stacknum
//		sum%=div
//
//	}
//	return sum%div
//}

//但是很难懂
//要结合动态规划的思想去看
func sumSubarrayMins(arr []int) int {
	//单调栈
	stack := []int{-1}
	sum := 0
	f := make([]int, len(arr))
	//以f(idx)表示以第idx个元素为右边界连续子数组的和，我们需要求f(4)
	//以4为右边界的连续子数组总共有4个
	//先假设4前面的元素是2,此时数组为[3,1,2,4]
	//此时f(3)表示以2为右边界的连续子数组的最小值之和
	//4>2，那么除了[4]以外的[2,4],[1,2,4]以及[3,1,2,4]三个连续子数组的和就是f(3)
	//假设4前面的元素是5，此时数组为[3,1,5,4]
	//4<5，那么除了[4],[4,5]以外的[1,5,4],[3,1,5,4]两个连续子数组的和就是f(2)
	div := int(math.Pow(10, 9) + 7)
	for k, v := range arr {
		//单调栈进出栈
		for len(stack) > 1 && v < arr[stack[len(stack)-1]] {
			//pop过程
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, k)
		f[k] = v * (k - stack[len(stack)-2])
		if stack[len(stack)-2] >= 0 {
			f[k] += f[stack[len(stack)-2]]
		}
		sum += f[k]
		sum %= div
	}
	return sum % div
}
