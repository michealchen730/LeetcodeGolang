package main

//根据字符串数组的长度从少到多排序
type Str01 []string

func (s Str01) Len() int           { return len(s) }
func (s Str01) Less(i, j int) bool { return len(s[i]) < len(s[j]) }
func (s Str01) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

//二维数组，坐标
//优先级：
//1.坐标第一位
//2.坐标第一位相同的情况下，使用坐标第二位
type Mat01 [][]int

func (m Mat01) Len() int { return len(m) }
func (m Mat01) Less(i, j int) bool {
	if m[i][0] < m[j][0] {
		return true
	} else if m[i][0] > m[j][0] {
		return false
	} else {
		return m[i][1] < m[j][1]
	}
}
func (m Mat01) Swap(i, j int) { m[i], m[j] = m[j], m[i] }

//按照字符串数组的长度排序
type StringsByLength []string

func (s StringsByLength) Len() int           { return len(s) }
func (s StringsByLength) Less(i, j int) bool { return len(s[i]) < len(s[j]) }
func (s StringsByLength) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
