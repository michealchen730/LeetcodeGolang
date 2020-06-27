package main

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
