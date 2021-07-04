package main

//func change(amount int, coins []int) int {
//	arr := make([]int, amount+1)
//	arr[0] = 1
//	sort.Ints(coins)
//
//	for _, v := range coins {
//		for i := 1; i <= amount; i++ {
//			if v <= i {
//				arr[i] += arr[i-v]
//			}
//		}
//	}
//	return arr[amount]
//}

//从背包问题的角度去考虑，不需要做排序，以及，j的循环可以从v开始
func change(amount int, coins []int) int {
	arr := make([]int, amount+1)
	arr[0] = 1
	for _, v := range coins {
		for j := v; j <= amount; j++ {
			arr[j] += arr[j-v]
		}
	}
	return arr[amount]
}
