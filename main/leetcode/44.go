package main

func isMatch2(s string, p string) bool {
	matrix := make([][]bool, len(p)+1)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]bool, len(s)+1)
	}
	matrix[0][0] = true
	for i := 1; i < len(matrix); i++ {
		if p[i-1] == '*' && matrix[i-1][0] {
			matrix[i][0] = true
		}
	}
	for i := 1; i < len(matrix); i++ {
	SI:
		for j := 1; j < len(matrix[0]); j++ {
			if p[i-1] != '*' {
				if (p[i-1] == s[j-1] || p[i-1] == '?') && matrix[i-1][j-1] {
					matrix[i][j] = true
				}
			}
			if p[i-1] == '*' {
				if matrix[i-1][j] || matrix[i][j-1] {
					for m := j; m < len(matrix[0]); m++ {
						matrix[i][m] = true
					}
					break SI
				}
			}
		}
	}
	return matrix[len(p)][len(s)]
}
