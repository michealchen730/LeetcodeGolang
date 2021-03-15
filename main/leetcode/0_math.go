package main

import "fmt"

func max(i int, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(x int, y int) int {
	if x > y {
		return y
	}
	return x
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//翻转int类型数组
func reverseArr(arr []int) {
	i, j := 0, len(arr)-1
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}

//KMP算法
func getNext(s string) []int {
	next := make([]int, len(s))
	next[0] = -1
	i, j := 0, -1
	for i < len(s)-1 {
		if j == -1 || s[i] == s[j] {
			i++
			j++
			next[i] = j
		} else {
			j = next[j]
		}
	}
	return next
}
func KMP(target string, text string) int {
	if len(text) == 0 {
		return 0
	}
	i, j := 0, 0
	next := getNext(text)
	for i < len(target) {
		if j == len(text)-1 && text[j] == target[i] {
			return i - j
		}
		if j == -1 || text[j] == target[i] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	return -1
}

//树算法
func getDepth(root *TreeNode) int {
	if root == nil {
		return 0
	} else {
		return max(getDepth(root.Left), getDepth(root.Right)) + 1
	}
}

//数学，i的二进制有多少位的长度
func lenBinary(i int) int {
	if i == 0 {
		return 0
	}
	res := 0
	for i != 0 {
		i >>= 1
		res++
	}
	return res
}

//求最大公约数
func gcd(x, y int) int {
	tmp := x % y
	if tmp > 0 {
		return gcd(y, tmp)
	} else {
		return y
	}
}

func arrToInt(arr []int) int {
	res := 0
	for i := 0; i < len(arr); i++ {
		res *= 10
		res += arr[i]
	}
	return res
}
func int2Bin(a int32) []int32 {
	var res []int32
	for b := a; b > 0; b = b / 2 {
		m := b % 2
		res = append(res, m)
	}
	return res
}

func forthBit(num int32) int32 {
	arr := int2Bin(num)
	fmt.Println(arr)
	if len(arr) < 4 {
		return 0
	} else {
		return arr[3]
	}
}

//最小堆
func segment(x int32, space []int) []int {
	if len(space) == 0 {
		return nil
	}
	var Q = make([]int, 0, len(space))
	res := make([]int, len(space)-int(x)+1)
	for i := 0; i < len(space); i++ {
		for len(Q) != 0 && space[i] < space[Q[len(Q)-1]] {
			Q = Q[:len(Q)-1]
		}
		Q = append(Q, i)
		if Q[0] == i-int(x) {
			Q = Q[1:]
		}
		if i+1-int(x) >= 0 {
			res[i+1-int(x)] = space[Q[0]]
		}
	}
	return res
}
