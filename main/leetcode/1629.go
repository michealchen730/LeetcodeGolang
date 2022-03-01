package main

func slowestKey(releaseTimes []int, keysPressed string) byte {
	mn := releaseTimes[0]
	res := keysPressed[0]
	for k := 1; k < len(releaseTimes); k++ {
		if releaseTimes[k]-releaseTimes[k-1] > mn || (releaseTimes[k]-releaseTimes[k-1] == mn && keysPressed[k] > res) {
			mn = releaseTimes[k] - releaseTimes[k-1]
			res = keysPressed[k]
		}
	}
	return res
}
