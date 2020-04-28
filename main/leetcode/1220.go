package main

func countVowelPermutation(n int) int {
	mat := []int{1, 1, 1, 1, 1}
	arr := []int{0, 0, 0, 0, 0}
	for i := 2; i <= n; i++ {
		arr[0] = (mat[1] + mat[2] + mat[4]) % 1000000007
		arr[1] = (mat[0] + mat[2]) % 1000000007
		arr[2] = (mat[1] + mat[3]) % 1000000007
		arr[3] = mat[2] % 1000000007
		arr[4] = (mat[2] + mat[3]) % 1000000007
		for k, v := range arr {
			mat[k] = v
		}
	}
	return (mat[0] + mat[1] + mat[2] + mat[3] + mat[4]) % 1000000007
}
