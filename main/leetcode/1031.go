package main

//前缀和暴力
func maxSumTwoNoOverlap(A []int, L int, M int) int {
	//前缀和
	arr := make([]int, len(A)+1)
	sum := 0
	for i := 0; i < len(A); i++ {
		sum += A[i]
		arr[i+1] = sum
	}
	mx := 0
	//处理,假设i点为当前L数组的结尾
	for i := L; i < len(arr); i++ {
		tp1 := arr[i] - arr[i-L]
		tp2 := 0
		//如果M数组在L数组左边
		for j := M; j <= i-L; j++ {
			tp2 = max(tp2, arr[j]-arr[j-M])
		}
		//如果M数组在L右边
		for j := len(arr) - 1; j >= i+M; j-- {
			tp2 = max(tp2, arr[j]-arr[j-M])
		}
		mx = max(mx, tp1+tp2)
	}
	return mx
}
