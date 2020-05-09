package main

func destCity(paths [][]string) string {
	mp1 := make(map[string]int)
	mp2 := make(map[string]int)
	for _, v := range paths {
		mp1[v[0]]++
		mp2[v[1]]++
	}
	for k, _ := range mp2 {
		if _, ok := mp1[k]; !ok {
			return k
		}
	}
	return ""
}
