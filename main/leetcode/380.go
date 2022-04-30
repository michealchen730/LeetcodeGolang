package main

import "math/rand"

//据说go的map结构在每次遍历的时候都会随机从某一个位置开始遍历
//但是这样是过不去这题的。。。
//type RandomizedSet struct {
//	mp map[int]int
//}
//
//
//func Constructor() RandomizedSet {
//	mp:=make(map[int]int)
//	return RandomizedSet{
//		mp: mp,
//	}
//}
//
//
//func (this *RandomizedSet) Insert(val int) bool {
//	if _,ok:=this.mp[val];ok{
//		return false
//	}
//	this.mp[val]=val
//	return true
//}
//
//
//func (this *RandomizedSet) Remove(val int) bool {
//	if _,ok:=this.mp[val];ok{
//		delete(this.mp,val)
//		return true
//	}
//	return false
//}
//
//
//func (this *RandomizedSet) GetRandom() int {
//	for k,_:=range this.mp{
//		return k
//	}
//	return -1
//}

type RandomizedSet struct {
	mp  map[int]int
	arr []int
}

func Constructor() RandomizedSet {
	mp := make(map[int]int)
	arr := make([]int, 0)
	return RandomizedSet{
		mp:  mp,
		arr: arr,
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.mp[val]; ok {
		return false
	}
	this.mp[val] = len(this.arr)
	this.arr = append(this.arr, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if idx, ok := this.mp[val]; ok {
		delete(this.mp, val)
		//最后一位移到删除值的位置上
		//如果移除的恰好是最后一位，那么直接移除即可！！！
		if idx != len(this.arr)-1 {
			this.arr[idx] = this.arr[len(this.arr)-1]
			//mp的赋值跟着改变
			this.mp[this.arr[idx]] = idx
		}
		//删除this.arr的最后一位
		this.arr = this.arr[:len(this.arr)-1]
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	return this.arr[rand.Intn(len(this.arr))]
}
