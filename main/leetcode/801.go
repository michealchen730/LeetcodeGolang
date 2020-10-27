package main

//func minSwap(A []int, B []int) int {
//	dp:=make([][]int,len(A)+1)
//	for k,_:=range dp{
//		dp[k]=make([]int,2)
//	}
//	dp[1][1]=1
//	for i:=2;i<len(dp);i++{
//		dp[i][0],dp[i][1]=2*len(A),2*len(A)
//		//表示当前位置不交换是合法的
//		if A[i-1]>A[i-2]&&B[i-1]>B[i-2]{
//			//如果不交换
//			dp[i][0]=min(dp[i-1][0],dp[i][0])
//			//如果交换，那么在上一位置交换的前提下，必合法
//			dp[i][1]=min(dp[i-1][1]+1,dp[i][1])
//
//			//如果前一位置交换也是合法的，那么也可以选择交换
//			if A[i-1]>B[i-2]&&B[i-1]>A[i-2]{
//				dp[i][0]=min(dp[i-1][1],dp[i][0])
//				dp[i][1]=min(dp[i][1],dp[i-1][0]+1)
//			}
//
//		}else{
//			//如果非法
//			//不交换当前位置，交换它的前一位置
//			dp[i][0]=min(dp[i][0],dp[i-1][1])
//			//或者交换当前位置，保持前一位置不变
//			dp[i][1]=min(dp[i][1],dp[i-1][0]+1)
//		}
//	}
//	return min(dp[len(A)][0],dp[len(A)][1])
//}

//两轮只需要用两个值
func minSwap(A []int, B []int) int {
	last0, last1 := 0, 1
	temp0, temp1 := 2*len(A), 2*len(A)
	for i := 2; i <= len(A); i++ {
		temp0, temp1 = 2*len(A), 2*len(A)
		//表示当前位置不交换是合法的
		if A[i-1] > A[i-2] && B[i-1] > B[i-2] {
			//如果不交换
			temp0 = min(last0, temp0)
			//如果交换，那么在上一位置交换的前提下，必合法
			temp1 = min(last1+1, temp1)

			//如果前一位置交换也是合法的，那么也可以选择交换
			if A[i-1] > B[i-2] && B[i-1] > A[i-2] {
				temp0 = min(last1, temp0)
				temp1 = min(temp1, last0+1)
			}

		} else {
			//如果非法
			//不交换当前位置，交换它的前一位置
			temp0 = min(temp0, last1)
			//或者交换当前位置，保持前一位置不变
			temp1 = min(temp1, last0+1)
		}
		last0, last1 = temp0, temp1
	}
	return min(temp0, temp1)
}
