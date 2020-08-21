package main

import (
	"errors"
	"fmt"
)

//自平衡的二叉搜索树实现
type AVLNode struct {
	Left, Right *AVLNode    // 表示指向左孩子和右孩子
	Data        interface{} // 结点存储数据
	height      int         // 记录这个结点此时的高度
}

// 定义comparer 指针类型
// 用来比较两个结点中Data的大小
type comparator func(a, b interface{}) int

// compare 指针
var compare comparator

// 新建一个结点
func NewAVLNode(data interface{}) *AVLNode {
	return &AVLNode{
		Left:   nil,
		Right:  nil,
		Data:   data,
		height: 1,
	}
}

// 新建AVL 树
func NewAVLTree(data interface{}, myfunc comparator) (*AVLNode, error) {
	if data == nil && myfunc == nil {
		return nil, errors.New("不能为空")
	}
	compare = myfunc
	return NewAVLNode(data), nil
}

// 获取结点数据
func (avlNode *AVLNode) GetData() interface{} {
	return avlNode.Data
}

// 设置结点数据
func (avlNode *AVLNode) SetData(data interface{}) {
	if avlNode == nil {
		return
	}
	avlNode.Data = data
}

// 获取结点的右孩子结点
func (avlNode *AVLNode) GetRight() *AVLNode {
	if avlNode == nil {
		return nil
	}
	return avlNode.Right
}

// 获取结点的左孩子结点
func (avlNode *AVLNode) GetLeft() *AVLNode {
	if avlNode == nil {
		return nil
	}
	return avlNode.Left
}

// 获取结点的高度
func (avlNode *AVLNode) GetHeight() int {
	if avlNode == nil {
		return 0
	}
	return avlNode.height
}

//比较两个子树高度的大小
func Max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

// 查找指定值
func (avlNode *AVLNode) Find(data interface{}) *AVLNode {
	var finded *AVLNode
	// 调用比较函数比较两个结点的指的大小
	switch compare(data, avlNode.Data) {
	case -1:
		finded = avlNode.Left.Find(data)
	case 1:
		finded = avlNode.Right.Find(data)
	case 0:
		return avlNode
	}

	return finded
}

// 查找最大值
func (avlNode *AVLNode) FindMax() *AVLNode {
	var finded *AVLNode

	if avlNode.Right != nil {
		finded = avlNode.Right.FindMin()
	} else {
		finded = avlNode
	}

	return finded
}

// 查找最小值
func (avlNode *AVLNode) FindMin() *AVLNode { // 递归写法
	var finded *AVLNode

	if avlNode.Left != nil {
		finded = avlNode.Left.FindMin()
	} else {
		finded = avlNode
	}

	return finded
}

// 插入数据
// 因为没有定义 结点的parent指针，所以你插入数据就只能递归的插入，因为我要调整树的平衡和高度
func (avlNode *AVLNode) Insert(value interface{}) *AVLNode {
	if avlNode == nil {
		return NewAVLNode(value)
	}

	switch compare(value, avlNode.Data) {
	case -1:
		// 如果value 小于 avlNode.Data 那么就在avlNode的左子树上去插入value
		avlNode.Left = avlNode.Left.Insert(value)
		avlNode = avlNode.adjust() // 自动调整平衡
	case 1:
		avlNode.Right = avlNode.Right.Insert(value)
		avlNode = avlNode.adjust()
	case 0:
		fmt.Print("数据已经存在")
	}
	// 修改结点的高度
	avlNode.height = Max(avlNode.Left.GetHeight(), avlNode.Right.GetHeight()) + 1

	return avlNode
}

// 删除数据
func (avlNode *AVLNode) Delete(value interface{}) *AVLNode {
	// 没有找到匹配的数据
	if avlNode == nil {
		//fmt.Println("不存在", value)
		return nil
	}

	switch compare(value, avlNode.Data) {
	case 1:
		avlNode.Right = avlNode.Right.Delete(value)
	case -1:
		avlNode.Left = avlNode.Left.Delete(value)
	case 0:
		// 找到数据，删除结点
		if avlNode.Left != nil && avlNode.Right != nil { // 结点有左孩子和右孩子
			avlNode.Data = avlNode.Right.FindMin().Data
			avlNode.Right = avlNode.Right.Delete(avlNode.Data)
		} else if avlNode.Left != nil { // 结点只有左孩子，无右孩子
			avlNode = avlNode.Left
		} else { // 结点只有右孩子或者无孩子
			avlNode = avlNode.Right
		}

	}

	// 自动调整平衡, 如果avlNode!=nil说明执行了对avlNode 的某个子树执行了删除结点，那么就需要重新调整树的平衡
	if avlNode != nil {
		avlNode.height = Max(avlNode.Left.GetHeight(), avlNode.Right.GetHeight()) + 1
		avlNode = avlNode.adjust() // 自动平衡
	}

	return avlNode
}

// 调整AVL树的平衡
func (avlNode *AVLNode) adjust() *AVLNode {
	// 如果右子树的高度比左子树的高度大于2
	if avlNode.Right.GetHeight()-avlNode.Left.GetHeight() == 2 {
		// 如果 avlNode.Right 的右子树的高度比avlNode.Right的左子树高度大
		// 直接对avlNode进行左旋转
		// 否则先对 avlNode.Right进行右旋转然后再对avlNode进行左旋转
		if avlNode.Right.Right.GetHeight() > avlNode.Right.Left.GetHeight() {
			avlNode = avlNode.LeftRotate()
		} else {
			avlNode = avlNode.RightThenLeftRotate()
		}
		// 如果左子树的高度比右子树的高度大2
	} else if avlNode.Right.GetHeight()-avlNode.Left.GetHeight() == -2 {
		// 如果avlNode.Left的左子树高度大于avlNode.Left的右子树高度
		// 那么就直接对avlNode进行右旋
		// 否则先对avlNode.Left进行左旋，然后对avlNode进行右旋
		if avlNode.Left.Left.GetHeight() > avlNode.Left.Right.GetHeight() {
			avlNode = avlNode.RightRotate()
		} else {
			avlNode = avlNode.LeftThenRightRotate()
		}
	}

	return avlNode
}

// 先顺时针旋转再逆时针旋转，先右旋，再左旋
func (avlNode *AVLNode) RightThenLeftRotate() *AVLNode {
	// 先把右孩子进行右旋
	avlNode.Right = avlNode.Right.RightRotate()
	// 然后把自己右旋
	return avlNode.LeftRotate()
}

// 逆时针旋转，左旋
func (avlNode *AVLNode) LeftRotate() *AVLNode {
	headNode := avlNode.Right
	avlNode.Right = headNode.Left
	headNode.Left = avlNode

	// 更新结点的高度
	// 这里应该注意的俄式应该先更新avlNode 的高度，因为headNode结点在avlNode结点的上面
	// headNode计算高度的时候要根据avlNode的高度来计算
	avlNode.height = Max(avlNode.Left.GetHeight(), avlNode.Right.GetHeight()) + 1
	headNode.height = Max(headNode.Left.GetHeight(), headNode.Right.GetHeight()) + 1

	return headNode
}

// 先逆时针旋转再顺时针旋转，先左旋，在右旋
func (avlNode *AVLNode) LeftThenRightRotate() *AVLNode {
	// 先把左孩子结点进行左旋
	avlNode.Left = avlNode.Left.LeftRotate()
	// 然后把自己右旋
	return avlNode.RightRotate()
}

// 顺时针旋转，右旋
func (avlNode *AVLNode) RightRotate() *AVLNode {
	headNode := avlNode.Left
	avlNode.Left = headNode.Right
	headNode.Right = avlNode

	// 更新旋转后结点的高度
	avlNode.height = Max(avlNode.Left.GetHeight(), avlNode.Right.GetHeight()) + 1
	headNode.height = Max(headNode.Left.GetHeight(), headNode.Right.GetHeight()) + 1

	return headNode
}
