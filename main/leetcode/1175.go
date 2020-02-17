package main

func numPrimeArrangements(n int) int {
	if n<=2{
		return 1
	}
	primes:=[]int{2,3,5,7,11,13,17,19,23,29,31,37,41,43,47,53,59,61,67,71,73,79,83,89,97,101}
	nmp:=0
	for primes[nmp]<=n{
		nmp++
	}
	//此时得到的nmp为质数的个数
	nmn:=n-nmp
	short:=0
	long:=0
	if nmp>nmn {
		short,long=nmn,nmp
	}else{
		short,long=nmp,nmn
	}
	sum:=1
	for i:=1;i<=long;i++{
		if sum>=1000000007 {
			sum=sum%1000000007
		}
		if i<=short {
			sum*=i
		}
		sum*=i
	}
	return sum%1000000007
}
