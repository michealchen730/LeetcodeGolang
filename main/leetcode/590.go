package main

//func postorder(root *Node) []int {
//	var res []int
//	if root!=nil{
//		queue:=[]*Node{root}
//		mp:=make(map[*Node]bool)
//		for len(queue)!=0{
//			tmpNode:=queue[len(queue)-1]
//			if mp[tmpNode]||len(tmpNode.Children)==0{
//				res=append(res,tmpNode.Val)
//				queue=queue[:len(queue)-1]
//			}else{
//				for i:=len(tmpNode.Children)-1;i>=0;i--{
//					queue=append(queue,tmpNode.Children[i])
//				}
//				mp[tmpNode]=true
//			}
//		}
//	}
//	return res
//}

func postorder(root *Node) []int {
	var res []int
	if root != nil {
		for _, v := range root.Children {
			res = append(res, postorder(v)...)
		}
		res = append(res, root.Val)
	}
	return res
}
