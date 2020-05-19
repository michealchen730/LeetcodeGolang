package main

import (
	"fmt"
	"sort"
)

type S937 []string

func (b S937) Len() int { return len(b) }
func (b S937) Less(i, j int) bool {
	m, n := 0, 0
	//找到第一个空白
	for ; m < len(b[i]); m++ {
		if b[i][m] == ' ' {
			break
		}
	}
	for ; n < len(b[j]); n++ {
		if b[j][n] == ' ' {
			break
		}
	}
	//判断空白符后的元素大小（如果都一样，会判定到最后）
	for m < len(b[i]) && n < len(b[j]) {
		if b[i][m] < b[j][n] {
			return true
		} else if b[i][m] > b[j][n] {
			return false
		} else {
			m++
			n++
		}
	}
	//判定到最后了，这时先看字符串长度，长度短的排在前面
	if len(b[i])-m != len(b[j])-n {
		return len(b[i])-m < len(b[j])-n
	} else {
		//说明长度也一样，这时候判断空白符前面的元素
		t := 0
		for t < len(b[i]) && t < len(b[j]) {
			if b[i][t] < b[j][t] {
				return true
			} else if b[i][t] > b[j][t] {
				return false
			} else {
				t++
			}
		}
		return true
	}
}
func (b S937) Swap(i, j int) { b[i], b[j] = b[j], b[i] }

func reorderLogFiles(logs []string) []string {
	var numlogs, wordlogs []string
	for _, v := range logs {
		for k := 0; k < len(v); k++ {
			if v[k] == ' ' {
				if v[k+1] >= 'a' {
					wordlogs = append(wordlogs, v)
				} else {
					numlogs = append(numlogs, v)
				}
				break
			}
		}
	}
	sort.Sort(S937(wordlogs))
	res := append(wordlogs, numlogs...)
	return res
}

func main() {
	fmt.Println(reorderLogFiles([]string{"a1 9 2 3 1", "g1 act car", "zo4 4 7", "ab1 off key dog", "a8 act zoo", "a2 act car"}))
}
