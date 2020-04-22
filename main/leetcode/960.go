package main

func minDeletionSize(A []string) int {
	dp := make([]int, len(A[0]))
	dp[0] = 1
	res := 1
	for k := 1; k < len(dp); k++ {
		//暴力的方式往回搜索（某种意义上的贪心原则）
		//这里也许可以考虑使用优先队列（完全意义上的贪心原则），存储当前dp中最大的列下标，顺着列下标往回搜索
		mx := 1
		for j := k - 1; j >= 0; j-- {
			//对于明确了小于当前mx的dp列，可以直接忽略
			//这算是一种优化，但我觉得优先队列会更高效
			if dp[j] <= mx-1 {
				continue
			}
			flag := 0
			for _, v := range A {
				if v[k] < v[j] {
					flag = 1
					break
				}
			}
			if flag == 0 && 1+dp[j] > mx {
				mx = 1 + dp[j]
			}
		}
		dp[k] = mx
		if mx > res {
			res = mx
		}
	}
	return len(A[0]) - res
}

//这个的思想应该是更快的，但实现起来复杂度反而增加了，优先队列的问题
func minDeletionSize2(A []string) int {
	dp := make([]int, len(A[0]))
	prior, res := []int{0}, 1
	dp[0] = 1
	for k := 1; k < len(dp); k++ {
		//使用优先队列（完全意义上的贪心原则），存储当前dp中最大的列下标，顺着列下标往回搜索
		mx := 1
	SC:
		for j := len(prior) - 1; j >= 0; j-- {
			//优先队列
			flag := 0
			for _, v := range A {
				if v[k] < v[prior[j]] {
					flag = 1
					break
				}
			}
			if flag == 0 && 1+dp[prior[j]] > mx {
				//如果找到满足条件的，那么根据优先级，此时的mx一定是能取到的最大的值
				mx = 1 + dp[prior[j]]
				break SC
			}
		}
		dp[k] = mx
		if mx > res {
			res = mx
		}
		pos := 0
		for pos < len(prior) && dp[prior[pos]] <= dp[k] {
			pos++
		}
		prior = append(prior[0:pos], append([]int{k}, prior[pos:]...)...)
	}
	return len(A[0]) - res
}
