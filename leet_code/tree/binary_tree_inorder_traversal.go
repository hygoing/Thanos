package main

import (
	. "persion.com/thanos/common"
	"fmt"
)

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	tempStore := make([]*TreeNode, 0)

	if root == nil {
		return res
	}

	node := root
	for node != nil || len(tempStore) > 0 {
		for node != nil {
			tempStore = append(tempStore, node)
			node = node.Left
		}
		node = tempStore[len(tempStore)-1]
		tempStore = tempStore[:len(tempStore)-1]
		res = append(res, node.Val)

		node = node.Right
	}
	return res
}


func main() {
	a := &TreeNode{
		Val: 1,
	}
	b := &TreeNode{
		Val: 2,
	}
	c := &TreeNode{
		Val: 3,
	}
	d := &TreeNode{
		Val: 4,
	}
	e := &TreeNode{
		Val: 5,
	}
	f := &TreeNode{
		Val: 6,
	}
	g := &TreeNode{
		Val: 7,
	}

	a.Left = b
	a.Right = c

	b.Left = d
	b.Right = e

	c.Left = f
	c.Right = g

	res := inorderTraversal(a)
	fmt.Println(res)
}
