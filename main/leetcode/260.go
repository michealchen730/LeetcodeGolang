package main

func singleNumber260(nums []int) []int {
	bitmask := 0
	for _, v := range nums {
		bitmask ^= v
	}
	diff := bitmask & (-bitmask)
	x := 0
	for _, v := range nums {
		if v&diff != 0 {
			x ^= v
		}
	}
	return []int{x, bitmask ^ x}
}
