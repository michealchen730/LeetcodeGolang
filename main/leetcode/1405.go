package main

import (
	"sort"
)

type LDS struct {
	Left   int
	B      byte
	Locked bool
}

func longestDiverseString(a int, b int, c int) string {
	lds := []LDS{
		{
			B:      'a',
			Left:   a,
			Locked: false,
		},
		{
			B:      'b',
			Left:   b,
			Locked: false,
		},
		{
			B:      'c',
			Left:   c,
			Locked: false,
		},
	}
	var res []byte
	//当abc有一个为负数，置locked
	//当再添加abc时会形成快乐字符串，置locked
	for true {
		//排序，找到应该添加的字母
		sort.Slice(lds, func(i, j int) bool {
			if lds[i].Locked {
				return false
			}
			if lds[j].Locked {
				return true
			}
			return lds[i].Left > lds[j].Left
		})
		//如果该字母锁定了（意味着所有字母都锁定）或者该字母剩余次数为0，说明字符串的构建结束了
		if lds[0].Locked == true || lds[0].Left == 0 {
			return string(res)
		} else {
			//如果使用该字母会形成快乐字符串，锁定该字母
			if len(res) > 1 && res[len(res)-1] == lds[0].B && res[len(res)-2] == lds[0].B {
				lds[0].Locked = true
				continue
			}
			//否则使用该字母
			res = append(res, lds[0].B)
			lds[0].Left--
			lds[1].Locked = false
			lds[2].Locked = false
		}
	}
	return string(res)
}
