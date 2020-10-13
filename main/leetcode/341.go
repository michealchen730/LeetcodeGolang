package main

//type NestedInteger struct {
//}
//
//func (this NestedInteger) IsInteger() bool {}
//
//func (this NestedInteger) GetInteger() int {}
//
//func (n *NestedInteger) SetInteger(value int) {}
//
//func (this *NestedInteger) Add(elem NestedInteger) {}
//
//func (this NestedInteger) GetList() []*NestedInteger {}
//
//
//
//
//type NestedIterator struct {
//	list []int
//	size int
//	temp int
//}
//
//func Constructor(nestedList []*NestedInteger) *NestedIterator {
//	var lis []int
//	for i:=0;i<len(nestedList);i++{
//		lis=append(lis,readNestedInteger(*nestedList[i])...)
//	}
//	return &NestedIterator{list: lis,size: len(lis),temp: 0}
//}
//
//func (this *NestedIterator) Next() int {
//	t:=this.temp
//	this.temp++
//	return this.list[t]
//}
//
//func (this *NestedIterator) HasNext() bool {
//	if this.temp==this.size{
//		return false
//	}else{
//		return true
//	}
//}
//
//func readNestedInteger(n NestedInteger) []int{
//	var res []int
//	if n.IsInteger(){
//		return []int{n.GetInteger()}
//	}else{
//		list:=n.GetList()
//		for _,v:=range list{
//			res=append(res,readNestedInteger(*v)...)
//		}
//	}
//	return res
//}
