package main

func oneEditAway(first string, second string) bool {
	f, s := len(first), len(second)
	if abs(f-s) > 1 {
		return false
	} else if f == s {
		t := 0
		for i := 0; i < f; i++ {
			if first[i] != second[i] {
				t++
			}
			if t > 1 {
				return false
			}
		}
		return true
	} else {
		mat := make([][]int, f+1)
		for k, _ := range mat {
			mat[k] = make([]int, s+1)
		}
		for i := 0; i < f; i++ {
			for j := 0; j < s; j++ {
				mat[i+1][j+1] = max(mat[i][j+1], mat[i+1][j])
				if first[i] == second[j] {
					mat[i+1][j+1] = max(1+mat[i][j], mat[i+1][j+1])
				}
			}
		}
		if mat[f][s] >= max(f, s)-1 {
			return true
		} else {
			return false
		}
	}
}
