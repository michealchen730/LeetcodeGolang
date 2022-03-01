package main

func isOneBitCharacter(bits []int) bool {
	temp := 0
	for temp < len(bits) {
		if temp == len(bits)-1 && bits[temp] == 0 {
			return true
		}
		if bits[temp] == 1 {
			temp++
		}
		temp++
	}
	return false
}
