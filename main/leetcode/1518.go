package main

func numWaterBottles(numBottles int, numExchange int) int {
	res := 0
	for numBottles >= numExchange {
		t := numBottles % numExchange
		res += numBottles - t
		numBottles = numBottles/numExchange + t
	}
	res += numBottles
	return res
}
