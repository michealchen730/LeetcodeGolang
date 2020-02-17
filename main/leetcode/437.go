package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}



func pathSum(root *TreeNode, sum int) int {
	if root==nil {
		return 0
	}
	var nodes []TreeNode
	traverse(root,&nodes)
	res:=0
	for _,v:=range nodes{
		findKinTree(&v,0,sum,&res)
	}
	return res
}

func traverse(root *TreeNode,nodes *[]TreeNode) {
	if root!=nil {
		*nodes=append(*nodes, *root)
		if root.Left!=nil {
			traverse(root.Left,nodes)
		}
		if root.Right!=nil {
			traverse(root.Right,nodes)
		}
	}
}

func findKinTree(root *TreeNode,sum int,k int,res *int)  {
	if root!=nil {
		sum=sum+root.Val
		if sum==k {
			*res++
		}
		if root.Left!=nil {
			findKinTree(root.Left,sum,k,res)
		}
		if root.Right!=nil {
			findKinTree(root.Right,sum,k,res)
		}
	}
}
