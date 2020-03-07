package main

func longestCommonSubsequence(text1 string, text2 string) int {
	arr := make([][]int, len(text1)+1)
	for i := 0; i < len(arr); i++ {
		arr[i] = make([]int, len(text2)+1)
	}
	for i := 1; i < len(arr); i++ {
		for j := 1; j < len(arr[0]); j++ {
			if text1[i-1] == text2[j-1] {
				arr[i][j] = arr[i-1][j-1] + 1
			} else {
				arr[i][j] = max(arr[i-1][j], arr[i][j-1])
			}
		}
	}
	return arr[len(text1)][len(text2)]
}
