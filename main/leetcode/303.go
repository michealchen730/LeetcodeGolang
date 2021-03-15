package main

type NumArray struct {
	arr []int
}

func Constructor303(nums []int) NumArray {
	arr := make([]int, len(nums)+1)
	for k, v := range nums {
		arr[k+1] = arr[k] + v
	}
	return NumArray{
		arr: arr,
	}
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.arr[j+1] - this.arr[i]
}
