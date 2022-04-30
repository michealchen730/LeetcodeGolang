package main

//func preorder(root *Node) []int {
//	var res []int
//	if root!=nil{
//		res=append(res,root.Val)
//		for _,v:=range root.Children{
//			res=append(res,preorder(v)...)
//		}
//	}
//	return res
//}

//func preorder(root *Node) []int {
//	var res []int
//	if root!=nil{
//		queue:=[]*Node{root}
//		for len(queue)!=0{
//			tmpNode:=queue[0]
//			//在这一步上会增加很多消耗
//			queue=append(tmpNode.Children,queue[1:]...)
//			res=append(res,tmpNode.Val)
//		}
//	}
//	return res
//}

func preorder(root *Node) []int {
	var res []int
	if root != nil {
		queue := []*Node{root}
		for len(queue) != 0 {
			tmpNode := queue[len(queue)-1]
			queue = queue[:len(queue)-1]
			res = append(res, tmpNode.Val)
			for i := len(tmpNode.Children) - 1; i >= 0; i-- {
				queue = append(queue, tmpNode.Children[i])
			}
		}
	}
	return res
}
