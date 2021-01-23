package main

import "fmt"

//01背包问题
//二维dp
//CW[i]表示所有可选的物品，CW[i][0]表示Cost，CW[i][1]表示Worth
//V表示背包的容积
//func ZeroOnePack(CW [][]int,V int) int {
//	N:=len(CW)
//	dp:=make([][]int,N+1)
//	for k,_:=range dp{
//		dp[k]=make([]int,V+1)
//	}
//	//dp[i][j]表示前i个物品放入容器为j的背包中的最大价值
//	for i:=1;i<=N;i++{
//		C,W:=CW[i-1][0],CW[i-1][1]
//		for j:=1;j<=V;j++{
//			dp[i][j]=dp[i-1][j]
//			if j-C>=0{
//				dp[i][j]=max(dp[i][j],dp[i-1][j-C]+W)
//			}
//		}
//	}
//	return dp[N][V]
//}

//一维dp，将V逆序遍历即可
func ZeroOnePack(CW [][]int, V int) int {
	N := len(CW)
	dp := make([]int, V+1)

	for i := 1; i <= N; i++ {
		C, W := CW[i-1][0], CW[i-1][1]
		for j := V; j >= 1; j-- {
			if j-C >= 0 {
				dp[j] = max(dp[j], dp[j-C]+W)
			}
		}
	}
	return dp[V]
}

//完全背包问题
//一维dp，将V逆序遍历即可
func CompletePack(CW [][]int, V int) int {
	N := len(CW)
	dp := make([]int, V+1)

	for i := 1; i <= N; i++ {
		C, W := CW[i-1][0], CW[i-1][1]
		for j := 1; j <= V; j++ {
			if j-C >= 0 {
				dp[j] = max(dp[j], dp[j-C]+W)
			}
		}
	}
	return dp[V]
}

//多重背包通用型解法
func MultiplePack(CW [][]int, V int) int {
	N := len(CW)
	dp := make([]int, V+1)
	for i := 1; i <= N; i++ {
		C, W, cnt := CW[i-1][0], CW[i-1][1], CW[i-1][2]
		for j := V; j >= 1; j-- {
			for c := 1; c <= cnt; c++ {
				if j-C*c >= 0 {
					dp[j] = max(dp[j], dp[j-C*c]+W*c)
				} else {
					break
				}
			}
		}
	}
	return dp[V]
}

//多重背包二进制问题
func MultiplePack2(CW [][]int, V int) int {
	//处理背包
	var Pack01 [][]int
	for _, v := range CW {
		vi, wi, si := v[0], v[1], v[2]
		temp := 1
		for si >= temp {
			Pack01 = append(Pack01, []int{temp * vi, temp * wi})
			si -= temp
			temp *= 2

		}
		if si > 0 {
			Pack01 = append(Pack01, []int{si * vi, si * wi})
		}
	}
	fmt.Println(Pack01)
	return ZeroOnePack(Pack01, V)
}

//func main(){
//	CW:=[][]int{[]int{1,2,3},[]int{2,4,1},[]int{3,4,3},[]int{4,5,2}}
//	fmt.Println(MultiplePack2(CW,5))
//}
