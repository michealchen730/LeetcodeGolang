package main

//模拟会导致超时
//func minKBitFlips(A []int, K int) int {
//	res := 0
//	for i:=0;i<len(A);i++{
//		if i+K>len(A)&&A[i]==0{
//			return -1
//		}
//		if A[i]==0{
//			res++
//			for j:=1;j<i+K;j++{
//				if A[j]==0{
//					A[j]++
//				}else{
//					A[j]--
//				}
//			}
//		}
//	}
//	return res
//}

func minKBitFlips(A []int, K int) int {
	ans := 0
	n := len(A)
	revCnt := 0
	for i, v := range A {
		//因为对每个i来说，它前面无论反转了奇数次还是偶数次，最终都能归结到1次或者0次
		//但无论是1次或是0次，需要判断一下是否会导致当前的这个i反转
		//比如0001，轮到1的时候，它前面反转了1次，但这个一次并不会影响到i的值
		//因此，只需要看到i-K个元素是否反转，就能知道i元素前面的反转是否会影响到第i个元素
		if i >= K && A[i-K] > 1 {
			revCnt ^= 1
			A[i-K] -= 2 // 复原数组元素，若允许修改数组 A，则可以省略
		}
		//在K范围内revCnt可以直接用
		if v == revCnt {
			if i+K > n {
				return -1
			}
			ans++
			revCnt ^= 1
			A[i] += 2
		}
	}
	return ans
}
