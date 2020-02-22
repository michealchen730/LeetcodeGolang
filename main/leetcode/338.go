package main

func countBits(num int) []int {
	if num == 0 {
		return []int{0}
	}
	if num == 1 {
		return []int{0, 1}
	}
	arr := make([]int, num+1)
	arr[0], arr[1], arr[2] = 0, 1, 1
	i := 2
	for arr[num] == 0 {
		temp, j := i, i+1
		i *= 2
		for j < i && j <= num {
			arr[j] = arr[j-temp] + arr[temp]
			j++
		}
		if i <= num {
			arr[i] = 1
		}
	}
	return arr
}
