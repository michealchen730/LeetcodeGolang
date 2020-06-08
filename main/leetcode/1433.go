package main

import (
	"fmt"
	"sort"
)

func checkIfCanBreak(s1 string, s2 string) bool {
	b1 := []byte(s1)
	b2 := []byte(s2)
	sort.Sort(Bytes1079(b1))
	sort.Sort(Bytes1079(b2))
	mn := 27
	mx := -27
	for i := 0; i < len(b1); i++ {
		mn = min(int(b1[i])-int(b2[i]), mn)
		mx = max(int(b1[i])-int(b2[i]), mx)
		if mn < 0 && mx > 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(checkIfCanBreak("abe", "acd"))
}
