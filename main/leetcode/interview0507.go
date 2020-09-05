package main

func exchangeBits(num int) int {
	odd := num & 0x55555555
	even := num & 0xaaaaaaaa
	odd = odd << 1
	even = even >> 1
	return odd | even
}
