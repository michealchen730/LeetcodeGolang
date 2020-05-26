package main

type MapSum struct {
	mp map[string]int
}

/** Initialize your data structure here. */
func Constructor677() MapSum {
	return MapSum{mp: make(map[string]int)}
}

func (this *MapSum) Insert(key string, val int) {
	this.mp[key] = val
}

func (this *MapSum) Sum(prefix string) int {
	res := 0
	for k, v := range this.mp {
		if len(k) >= len(prefix) && k[0:len(prefix)] == prefix {
			res += v
		}
	}
	return res
}
