package main

//func findAnagrams(s string, p string) []int {
//	if len(s)<len(p){
//		return nil
//	}
//	//记录p的数值
//	mp:=make(map[byte]int)
//	for k := range p{
//		mp[p[k]]++
//	}
//	//初始化滑动窗口
//	mp2:=make(map[byte]int)
//	for i:=0;i<len(p);i++{
//		mp2[s[i]]++
//	}
//	var res []int
//	if check438(mp2,mp){
//		res=append(res,0)
//	}
//	for i:=len(p);i<len(s);i++{
//		mp2[s[i]]++
//		mp2[s[i-len(p)]]--
//		if check438(mp2,mp){
//			res=append(res,i-len(p)+1)
//		}
//	}
//	return res
//}
//
//func check438(mp1 map[byte]int, mp2 map[byte]int) bool{
//	for k,v:=range mp2{
//		if mp1[k]!=v{
//			return false
//		}
//	}
//	return true
//}

// 先看下将第二个map转数组是否会提速
// 从36ms提到了12ms 50%
//func findAnagrams(s string, p string) []int {
//	if len(s)<len(p){
//		return nil
//	}
//	//记录p的数值
//	mp:=make(map[byte]int)
//	for k := range p{
//		mp[p[k]]++
//	}
//	//初始化滑动窗口
//	mp2:=make([]int,26)
//	for i:=0;i<len(p);i++{
//		mp2[s[i]-'a']++
//	}
//	var res []int
//	if check438(mp2,mp){
//		res=append(res,0)
//	}
//	for i:=len(p);i<len(s);i++{
//		mp2[s[i]-'a']++
//		mp2[s[i-len(p)]-'a']--
//		if check438(mp2,mp){
//			res=append(res,i-len(p)+1)
//		}
//	}
//	return res
//}
//
//func check438(mp1 []int, mp2 map[byte]int) bool{
//	for k,v:=range mp2{
//		if mp1[k-'a']!=v{
//			return false
//		}
//	}
//	return true
//}

//把第一个map改为数组
//速度提到了4ms 97%
//func findAnagrams(s string, p string) []int {
//	if len(s)<len(p){
//		return nil
//	}
//	//记录p的数值
//	mp:=make([]int,26)
//	for k := range p{
//		mp[p[k]-'a']++
//	}
//	//初始化滑动窗口
//	mp2:=make([]int,26)
//	for i:=0;i<len(p);i++{
//		mp2[s[i]-'a']++
//	}
//	var res []int
//	if check438(mp2,mp){
//		res=append(res,0)
//	}
//	for i:=len(p);i<len(s);i++{
//		mp2[s[i]-'a']++
//		mp2[s[i-len(p)]-'a']--
//		if check438(mp2,mp){
//			res=append(res,i-len(p)+1)
//		}
//	}
//	return res
//}
//
//func check438(mp1 []int, mp2 []int) bool{
//	for k,v:=range mp2{
//		if mp1[k]!=v{
//			return false
//		}
//	}
//	return true
//}

// 不是每次都要做check438的
// 但是先判断mp2[s[i]-'a']!=mp[s[i]-'a'] ||mp2[s[i-len(p)]-'a']!=mp[s[i-len(p)]-'a']是否命中再check438
// 这个整体流程未必比直接check438快多少
func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return nil
	}
	//记录p的数值
	mp := make([]int, 26)
	for k := range p {
		mp[p[k]-'a']++
	}
	//初始化滑动窗口
	mp2 := make([]int, 26)
	for i := 0; i < len(p); i++ {
		mp2[s[i]-'a']++
	}
	var res []int
	if check438(mp2, mp) {
		res = append(res, 0)
	}
	for i := len(p); i < len(s); i++ {
		mp2[s[i]-'a']++
		mp2[s[i-len(p)]-'a']--
		if mp2[s[i]-'a'] != mp[s[i]-'a'] || mp2[s[i-len(p)]-'a'] != mp[s[i-len(p)]-'a'] {
			continue
		}
		if check438(mp2, mp) {
			res = append(res, i-len(p)+1)
		}
	}
	return res
}

func check438(mp1 []int, mp2 []int) bool {
	for k, v := range mp2 {
		if mp1[k] != v {
			return false
		}
	}
	return true
}
