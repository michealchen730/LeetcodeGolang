package main

func isValid(s string) bool {
	if len(s)%2!=0{
		return false
	}
	var res []byte
	for i:=0;i<len(s);i++{
		switch s[i] {
		case '}':
			if len(res)!=0&&res[len(res)-1]=='{' {
				res=res[:len(res)-1]
			}else {
				return false
			}
		case ']':
			if len(res)!=0&&res[len(res)-1]=='[' {
				res=res[:len(res)-1]
			}else {
				return false
			}
		case ')':
			if len(res)!=0&&res[len(res)-1]=='(' {
				res=res[:len(res)-1]
			}else {
				return false
			}
		default:
			res = append(res, s[i])
		}
	}
	if len(res)!=0{
		return false
	}else{
		return true
	}
}
