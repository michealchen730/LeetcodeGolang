package main

//type CBTInserter struct {
//	Root *TreeNode
//	Nodes []*TreeNode
//}
//
//
//func Constructor(root *TreeNode) CBTInserter {
//	//bfs遍历所有结点(层次遍历)
//	var res []*TreeNode
//	var queue []*TreeNode
//	res=append(res, root)
//	queue=append(queue,root)
//	for len(queue)!=0{
//		level:=len(queue)
//		for i:=0;i<level;i++{
//			if queue[i].Left!=nil{
//				res=append(res, queue[i].Left)
//				queue=append(queue, queue[i].Left)
//			}
//			if queue[i].Right!=nil{
//				res=append(res, queue[i].Right)
//				queue=append(queue,queue[i].Right)
//			}
//		}
//		queue=queue[level:]
//	}
//	return CBTInserter{
//		Root: root,
//		Nodes: res,
//	}
//}
//
//
//
//func (this *CBTInserter) Insert(v int) int {
//	node:=&TreeNode{Val: v}
//	this.Nodes=append(this.Nodes,node)
//	length:=len(this.Nodes)
//	parent:=this.Nodes[length/2-1]
//	if length%2==0{
//		parent.Left=node
//	}else{
//		parent.Right=node
//	}
//	return parent.Val
//}
//
//
//func (this *CBTInserter) Get_root() *TreeNode {
//	return this.Root
//}

type CBTInserter struct {
	Root *TreeNode
}

//func Constructor(root *TreeNode) CBTInserter {
//	return CBTInserter{Root: root}
//}

//func (this *CBTInserter) Insert(v int) int {
//	return insertCBT(this.Root,v)
//}

func (this *CBTInserter) Insert(v int) int {
	var res int
	node := &TreeNode{Val: v}
	//bfs遍历所有结点(层次遍历)
	var queue []*TreeNode
	queue = append(queue, this.Root)
EI:
	for len(queue) != 0 {
		level := len(queue)
		for i := 0; i < level; i++ {
			if queue[i].Left == nil {
				res = queue[i].Val
				queue[i].Left = node
				break EI
			}
			if queue[i].Right == nil {
				res = queue[i].Val
				queue[i].Right = node
				break EI
			}
			queue = append(queue, []*TreeNode{queue[i].Left, queue[i].Right}...)
		}
		queue = queue[level:]
	}
	return res
}

//func (this *CBTInserter) Insert(v int) int {
//	return insertCBT(this.Root,v)
//}

//func insertCBT(root *TreeNode,val int) int {
//	if root.Left==nil{
//		root.Left=&TreeNode{Val: val}
//		return root.Val
//	}
//	if root.Right==nil{
//		root.Right=&TreeNode{Val: val}
//		return root.Val
//	}
//	lh,rh:=0,0
//	temp:=root.Left
//	for temp!=nil{
//		lh++
//		temp=temp.Right
//	}
//	temp=root.Right
//	for temp!=nil{
//		rh++
//		temp=temp.Right
//	}
//	if lh==rh{
//		return insertCBT(root.Left,val)
//	}else{
//		return insertCBT(temp,val)
//	}
//}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.Root
}
