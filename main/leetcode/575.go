package main

func distributeCandies575(candyType []int) int {
	mp := make(map[int]int)
	for _, v := range candyType {
		mp[v]++
		if len(mp) >= len(candyType)/2 {
			return len(candyType) / 2
		}
	}
	return len(mp)
}
