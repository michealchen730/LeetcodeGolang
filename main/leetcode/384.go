package main

import "math/rand"

type Solution struct {
	src []int
}

func Constructor(nums []int) Solution {
	return Solution{src: nums}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.src
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	length := len(this.src)
	cpy := make([]int, length)
	shuffle := make([]int, length)
	copy(cpy, this.src)
	for i := 0; i < len(shuffle); i++ {
		temp := rand.Intn(len(cpy))
		shuffle[i] = cpy[temp]
		cpy = append(cpy[:temp], cpy[temp+1:]...)
	}
	return shuffle
}

//调用库函数
/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle2() []int {
	length := len(this.src)
	shuffle := make([]int, length)
	copy(shuffle, this.src)
	rand.Shuffle(length, func(i, j int) {
		shuffle[i], shuffle[j] = shuffle[j], shuffle[i]
	})
	return shuffle
}
