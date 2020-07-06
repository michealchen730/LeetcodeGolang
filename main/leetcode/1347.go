package main

func minSteps(s string, t string) int {
	arr1, arr2 := make([]int, 26), make([]int, 26)
	for k, v := range s {
		arr1[v-'a']++
		arr2[t[k]-'a']++
	}
	res := 0
	for k, v := range arr1 {
		res += abs(arr2[k] - v)
	}
	return res / 2
}
