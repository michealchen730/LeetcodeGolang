package main

func countCharacters(words []string, chars string) int {
	arr:=make([]int,26)
	for i:=0;i<len(chars);i++{
		arr[chars[i]-'a']++
	}
	res:=0
	for i:=0;i<len(words);i++{
		flag:=0
		temp:=len(words[i])-1
		for j:=0;j<len(words[i]);j++{
			arr[words[i][j]-'a']--
			if arr[words[i][j]-'a']<0 {
				flag=1
				temp=j
				break
			}
		}
		for j:=0;j<=temp;j++{
			arr[words[i][j]-'a']++
		}
		if flag==0 {
			res+=len(words[i])
		}
	}
	return res
}