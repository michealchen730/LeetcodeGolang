package main

func lastRemaining(n int, m int) int {
	res := 0
	for i := 2; i <= n; i++ {
		res = (res + m) % i
	}
	return res
}

//暴力模拟，超时
func lastRemaining2(n int, m int) int {
	arr := make([]int, n)
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}
	flag := 1
	temp := 0
	for len(arr) != 1 {
		flag++
		temp++
		if temp == len(arr) {
			temp = 0
		}
		if flag == m {
			//注意temp+1溢出
			if temp == len(arr)-1 {
				arr = arr[:len(arr)-1]
				temp = 0
			} else {
				arr = append(arr[0:temp], arr[temp+1:]...)
			}
			flag = 1
		}
	}
	return arr[0]
}
