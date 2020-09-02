package main

//这是一年前的写法，效率很低
//下一轮需要改进
func buildTree106(inorder []int, postorder []int) *TreeNode {
	l := len(inorder)

	if l == 0 {
		return nil
	} else {
		res := constructTree106(inorder, postorder, 0, l-1, 0, l-1)
		return res
	}
}

func constructTree106(in, post []int, ps, pe, is, ie int) *TreeNode {
	node := TreeNode{Val: post[pe]}
	if pe == ps {
		return &node
	}
	for i := is; i <= ie; i++ {
		if in[i] == post[pe] {
			if i < ie {
				node.Right = constructTree106(in, post, pe-(ie-i), pe-1, i+1, ie)
			}
			if i > is {
				node.Left = constructTree106(in, post, ps, pe-(ie-i+1), is, i-1)
			}
		}
	}
	return &node
}
