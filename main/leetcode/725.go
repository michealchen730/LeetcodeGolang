package main

func splitListToParts(root *ListNode, k int) []*ListNode {
	node1 := root
	var nodes []*ListNode
	for node1 != nil {
		nodes = append(nodes, node1)
		node1 = node1.Next
	}
	if len(nodes) < k {
		for i := 0; i < len(nodes); i++ {
			nodes[i].Next = nil
		}
		for len(nodes) < k {
			nodes = append(nodes, nil)
		}
		return nodes
	} else {
		nums1 := len(nodes) / k
		nums2 := len(nodes) % k
		nums3 := 0
		var res []*ListNode
		for i := 0; i < len(nodes); i += nums1 {
			if nums3 < nums2 {
				i += 1
			}
			if i != 0 {
				nodes[i-1].Next = nil
			}
			res = append(res, nodes[i])
		}
		return res
	}
}
