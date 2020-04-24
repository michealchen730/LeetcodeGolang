package main

func minDeletionSize944(A []string) int {
	res := 0
	if len(A) == 1 {
		return res
	}
	for i := 0; i < len(A[0]); i++ {
		for j := 1; j < len(A); j++ {
			if int32(A[j][i])-int32(A[j-1][i]) < 0 {
				res++
				break
			}
		}
	}
	return res
}
