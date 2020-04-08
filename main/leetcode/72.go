package main

func minDistance(word1 string, word2 string) int {
	if len(word1) == 0 || len(word2) == 0 {
		return max(len(word1), len(word2))
	}
	arr := make([][]int, len(word1)+1)
	for k, _ := range arr {
		arr[k] = make([]int, len(word2)+1)
	}
	for i := 0; i < len(arr); i++ {
		arr[i][0] = i
	}
	for i := 0; i < len(arr[0]); i++ {
		arr[0][i] = i
	}
	for i := 1; i < len(arr); i++ {
		for j := 1; j < len(arr[0]); j++ {
			temp := arr[i-1][j-1]
			if word1[i-1] != word2[j-1] {
				temp++
			}
			arr[i][j] = minof3(temp, arr[i-1][j]+1, arr[i][j-1]+1)
		}
	}
	return arr[len(word1)][len(word2)]
}
func minof3(i, j, k int) int {
	if i > j {
		i, j = j, i
	}
	if i < k {
		return i
	}
	return k
}
