package main

import (
	"regexp"
	"strconv"
)


func myAtoi(str string) int {
	reg:=regexp.MustCompile("^\\s*[+-]?0*\\d{1,10}")
	if reg!=nil{
		res:=reg.FindString(str)
		if res=="" {
			return 0
		}else{
			bytes:=[]byte(res)
			i:=0
			for i<len(bytes) {
				if bytes[i]==' '||bytes[i]=='+'{
					i++
				}else {
					break
				}
			}
			bytes=bytes[i:]
			bts:=string(bytes)
			result,_:=strconv.Atoi(bts)
			if -2147483648 >result{
				return -2147483648
			}
			if result>2147483647 {
				return 2147483647
			}

			return result
		}
	}

	return 0
}


func myAtoi2(str string) int {
	if len(str)==0 {
		return 0
	}
	flag1:=-1 //flag1表示有没有遇到正负号,遇到+为1，遇到-为0
	flag2:=-1//flag2表示空格前有没有遇到0
	//这一段代码会帮我们找到合乎规范的数字开头（非0开头）和不合乎规范的其他任意开头
	i:=0
	for i<len(str)&&((str[i]==' '&&flag2==-1)||((str[i]=='+'||str[i]=='-')&&flag1==-1)||str[i]=='0'){
		if str[i]=='+'{
			flag1=1
			flag2=0
		}
		if str[i]=='-'{
			flag1=0
			flag2=0
		}
		if str[i]=='0'{
			flag2=0
			if flag1==-1{
				flag1=2
			}
		}
		i++
	}
	if i==len(str){
		return 0
	}
	//我们判断这个开头是合乎规范的还是不合乎规范的
	if str[i]>'0'&&str[i]<='9' {
		//合乎规范
		var re []byte
		if flag1==0 {
			re=append(re, '-')
		}
		j:=i
		//找到第一个非数字，如果全是数字，那么只取11位
		for j<len(str)&&j-i<=10&&str[j]>='0'&&str[j]<='9' {
			re=append(re,str[j])
			j++
		}
		s:=string(re)
		res,_:=strconv.Atoi(s)
		if -2147483648 >res{
			return -2147483648
		}
		if res>2147483647 {
			return 2147483647
		}
		return res
	}else{
		return 0
	}
}
