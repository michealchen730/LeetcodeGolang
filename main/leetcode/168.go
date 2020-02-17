package main

func convertToTitle(n int) string {
	imap:=make(map[int]byte)
	for i:=1; i<=26; i++ {
		imap[i]=byte('A'+i-1)
	}
	if n==0 {
		return ""
	}
	var res []byte
	for{
		if n%26!=0 {
			res=append(res, imap[n%26])
			n=n/26
		}else{
			res=append(res, 'Z')
			n=n/26-1
		}
		if n==0 {
			break
		}
	}
	swapBytes(res,0,len(res)-1)
	result:=string(res)
	return result
}
