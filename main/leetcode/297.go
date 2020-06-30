package main

import (
	"strconv"
	"strings"
)

type Codec297 struct {
}

func Constructor297() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var res strings.Builder
	res.WriteByte('[')

	if root != nil {
		strs := []string{strconv.Itoa(root.Val)}
		queue := []*TreeNode{root}
		for len(queue) != 0 {
			l := len(queue)
			for i := 0; i < l; i++ {
				temp := queue[i]
				if temp.Left != nil {
					strs = append(strs, strconv.Itoa(temp.Left.Val))
					queue = append(queue, temp.Left)
				} else {
					strs = append(strs, "null")
				}
				if temp.Right != nil {
					strs = append(strs, strconv.Itoa(temp.Right.Val))
					queue = append(queue, temp.Right)
				} else {
					strs = append(strs, "null")
				}
			}
			queue = queue[l:]
		}
		for i := len(strs) - 1; i >= 0; i-- {
			if strs[i] != "null" {
				strs = strs[:i+1]
				break
			}
		}
		for k, v := range strs {
			res.WriteString(v)
			if k != len(strs)-1 {
				res.WriteByte(',')
			}
		}
	}
	res.WriteByte(']')
	return res.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	//去掉头尾的括号
	data = data[1 : len(data)-1]
	if len(data) == 0 {
		return nil
	}
	split := strings.Split(data, ",")

	val, _ := strconv.Atoi(split[0])
	res := &TreeNode{Val: val}
	queue := []*TreeNode{res}

	split = split[1:]
	for len(queue) != 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			node := queue[i]
			if len(split) > 0 {
				if split[0] != "null" {
					value, _ := strconv.Atoi(split[0])
					newNode := &TreeNode{Val: value}
					node.Left = newNode
					queue = append(queue, newNode)
				}
				split = split[1:]
			}
			if len(split) > 0 {
				if split[0] != "null" {
					value, _ := strconv.Atoi(split[0])
					newNode := &TreeNode{Val: value}
					node.Right = newNode
					queue = append(queue, newNode)
				}
				split = split[1:]
			}
		}
		queue = queue[l:]
	}
	return res
}
