package main

import "fmt"

func modifyString(s string) string {
	b := []byte(s)
	for i := 0; i < len(s); i++ {
		if s[i] == '?' {
			t1 := 'e'
			if i-1 >= 0 {
				t1 = int32(b[i-1])
			}
			t2 := 'f'
			if i+1 < len(s) {
				t2 = int32(b[i+1])
			}

			if 'a' != t1 && 'a' != t2 {
				b[i] = 'a'
				continue
			}
			if 'b' != t1 && 'b' != t2 {
				b[i] = 'b'
				continue
			}
			if 'c' != t1 && 'c' != t2 {
				b[i] = 'c'
			}
		}
	}
	return string(b)
}

func main() {
	fmt.Println(modifyString("??yw?ipkj?"))
}
