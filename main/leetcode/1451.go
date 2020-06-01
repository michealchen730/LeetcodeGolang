package main

import "strings"

func arrangeWords(text string) string {
	fields := strings.Fields(text)
	fields[0] = strings.ToLower(fields[0])
	for i := 1; i < len(fields); i++ {
		// 获得当前需要比较的元素值。
		tmp := fields[i]
		// 内层循环控制 比较 并 插入
		for j := i - 1; j >= 0; j-- {
			// mySlice[i] 需要插入的元素
			// mySlice[j] 需要比较的元素
			if len(tmp) < len(fields[j]) {
				// 如果插入的元素小，交换位置。将后边的元素与前边的元素互换
				fields[j+1] = fields[j]
				// 将前面的数设置为当前需要交换的数
				fields[j] = tmp
			} else {
				// 由于是已经排序好的，则不需要再次比较。
				break
			}
		}
	}
	var res strings.Builder
	res.WriteByte(byte(fields[0][0] - 32))
	res.WriteString(fields[0][1:len(fields[0])])
	for i := 1; i < len(fields); i++ {
		res.WriteByte(' ')
		res.WriteString(fields[i])
	}
	return res.String()
}
