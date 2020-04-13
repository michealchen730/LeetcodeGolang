package main

import "math"

func numOfWays(n int) int {
	lol, sad := 6, 6
	for i := 1; i < n; i++ {
		newlol := (lol*3 + sad*2) % int(math.Pow(10, 9)+7)
		newsad := (lol*2 + sad*2) % int(math.Pow(10, 9)+7)
		lol = newlol
		sad = newsad
	}
	return (lol + sad) % int(math.Pow(10, 9)+7)
}
