package main

func hammingWeightO15(num uint32) int {
	res, t := 0, uint32(1)
	for i := 0; i < 32; i++ {
		if num&t != 0 {
			res++
		}
		t = t << 1
	}
	return res
}
