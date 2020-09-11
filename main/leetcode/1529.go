package main

func minFlips(target string) int {
	cnt := 0
	for _, v := range target {
		state := '0'
		if cnt%2 != 0 {
			state = '1'
		}
		if state != v {
			cnt++
		}
	}
	return cnt
}
