package main

func countTriplets(arr []int) int {
	res := 0
	for i := 0; i < len(arr)-1; i++ {
		for k := i + 1; k < len(arr); k++ {
			a := arr[i]
			b := arr[k]
			for m := k - 1; m > i; m-- {
				b ^= arr[m]
			}
			if a == b {
				res++
			}
			for n := i + 1; n < k; n++ {
				a ^= arr[n]
				b ^= arr[n]
				if a == b {
					res++
				}
			}
		}
	}
	return res
}
