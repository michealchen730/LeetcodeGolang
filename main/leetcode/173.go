package main

type BSTIterator struct {
	nodes []int
	temp  int
}

func Constructor173(root *TreeNode) BSTIterator {
	return BSTIterator{
		temp:  0,
		nodes: morrisVal(root),
	}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	this.temp++
	return this.nodes[this.temp-1]
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return len(this.nodes) > this.temp
}
