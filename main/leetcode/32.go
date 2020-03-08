package main

//这里理解错题意了，求成了最长有效括号子序列的长度了
//func longestValidParentheses(s string) int {
//	arr:=make([][]int,len(s))
//	for k,_:=range arr{
//		arr[k]=make([]int,len(s))
//	}
//	for i:=len(s)-1;i>=0;i--{
//		for j:=i;j<len(s);j++{
//			if i!=j {
//				if j-i==1&&s[i]!=s[j]&&s[i]=='(' {
//					arr[i][j]=1
//					continue
//				}
//				temp:=0
//				for k:=i; k<j; k++ {
//					if arr[i][k]+arr[k+1][j]>temp {
//						temp=arr[i][k]+arr[k+1][j]
//					}
//				}
//				if s[i]=='('&&s[j]==')'&&temp<1+arr[i+1][j-1] {
//					temp=1+arr[i+1][j-1]
//				}
//				arr[i][j]=temp
//			}
//		}
//	}
//	fmt.Println(arr)
//	return arr[0][len(s)-1]*2
//}

func longestValidParentheses(s string) int {
	res := 0
	arr := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == ')' {
			if i-1 >= 0 {
				if s[i-1] == '(' {
					if i-2 >= 0 {
						arr[i] = arr[i-2] + 2
					} else {
						arr[i] = 2
					}
				} else {
					if i-arr[i-1]-1 >= 0 && arr[i-1] > 0 && s[i-arr[i-1]-1] == '(' {
						if i-arr[i-1]-2 >= 0 {
							arr[i] = arr[i-1] + 2 + arr[i-arr[i-1]-2]
						} else {
							arr[i] = arr[i-1] + 2
						}
					}
				}
				if arr[i] > res {
					res = arr[i]
				}
			}
		}
	}
	return res
}

//func longestValidParentheses(s string) int {
//	res:=0
//	arr:=make([][]int,len(s))
//	for k,_:=range arr{
//		arr[k]=make([]int,len(s))
//	}
//	for i:=len(s)-1;i>=0;i--{
//		for j:=i;j<len(s);j++{
//			if i!=j {
//				if s[i]=='('&&s[j]==')'&&(j-i)%2!=0{
//					if j-i==1 {
//						arr[i][j]=2
//					}else{
//						if arr[i+1][j-1]!=0 {
//							arr[i][j]=j-i+1
//						}else{
//							for k:=i+1;k<=j-2;k+=2  {
//								if arr[i][k]!=0&&arr[k+1][j]!=0{
//									arr[i][j]=j-i+1
//									break
//								}
//							}
//						}
//					}
//				}
//			}
//			if arr[i][j]>res {
//				res=arr[i][j]
//			}
//		}
//	}
//	return res
//}
