package main

import (
	"fmt"
	"strconv"
)

func spiralMatrixIII(R int, C int, r0 int, c0 int) [][]int {
	res := [][]int{{r0, c0}}
	N := max(max(R-r0, r0-0+1), max(C-c0, c0-0+1))
	for len(res) < R*C {
		for i := 1; i <= N; i++ {
			//找到起始点
			sr, sc := r0-i, c0+i
			if sc >= 0 && sc < C {
				for j := 0; j < 2*i; j++ {
					sr++
					if sr >= 0 && sr < R {
						res = append(res, []int{sr, sc})
					}
				}
			} else {
				sr += 2 * i
			}
			if sr >= 0 && sr < R {
				for j := 0; j < 2*i; j++ {
					sc--
					if sc >= 0 && sc < C {
						res = append(res, []int{sr, sc})
					}
				}
			} else {
				sc -= 2 * i
			}
			if sc >= 0 && sc < C {
				for j := 0; j < 2*i; j++ {
					sr--
					if sr >= 0 && sr < R {
						res = append(res, []int{sr, sc})
					}
				}
			} else {
				sr -= 2 * i
			}
			if sr >= 0 && sr < R {
				for j := 0; j < 2*i; j++ {
					sc++
					if sc >= 0 && sc < C {
						res = append(res, []int{sr, sc})
					}
				}
			} else {
				sc += 2 * i
			}
		}
	}
	return res
}

func main() {
	iii := spiralMatrixIII(100, 100, 82, 76)
	fmt.Println(len(iii))
	fmt.Println(iii)
	mp := make(map[string]int)
	for k, v := range iii {
		key := strconv.Itoa(v[0]) + " " + strconv.Itoa(v[1])
		_, ok := mp[key]
		if !ok {
			mp[key] = 1
		} else {
			fmt.Println(key)
			fmt.Println(k)
		}
	}

}
