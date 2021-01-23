package main

//确定A字符串是B字符串的子集
func isSubSet(A, B string) bool {
	if len(A) > len(B) {
		return false
	}
	i, j := 0, 0
	for j < len(B) {
		if A[i] == B[j] {
			i++
		}
		if i == len(A) {
			return true
		}
		j++
	}
	return false
}
