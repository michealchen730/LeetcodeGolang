package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length:=len(nums1)+len(nums2)
	s:=length/2+1
	i,j:=0,0
	var res []int
	for{
		if len(res)==s {
			break
		}else{
			if i<len(nums1)&&j<len(nums2) {
				if nums1[i]<nums2[j] {
					res=append(res, nums1[i])
					i++
				}else{
					res=append(res,nums2[j])
					j++
				}
			}else{
				if i==len(nums1) {
					res=append(res,nums2[j])
					j++
				}else{
					res=append(res,nums1[i])
					i++
				}
			}
		}
	}
	if length%2==0 {
		return (float64(res[len(res)-2])+float64(res[len(res)-1]))/2
	}else {
		return float64(res[len(res)-1])
	}
}
