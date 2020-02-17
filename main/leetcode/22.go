package main

//深度优先？+剪枝
func generateParenthesis(n int) []string {
	var res []string
	getResults(2*n,&res,"",0,n)
	return res
}
func getResults(n int,res *[]string,temp string,flag int,left int){
	if len(temp)==n{
		*res=append(*res, temp)
		return
	}
	if flag>=0&&left>=0{
		if left>=1 {
			getResults(n, res, temp+"(", flag+1,left-1)
		}
		getResults(n, res, temp+")", flag-1,left)
	}
}
