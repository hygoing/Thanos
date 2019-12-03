package main

import (
	. "persion.com/thanos/common"
	"fmt"
)

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	tempStore := make([]*TreeNode, 0)

	if root == nil {
		return res
	}

	tempStore = append(tempStore, root)
	for len(tempStore) > 0 {
		count := len(tempStore)
		tempRes := make([]int, 0)
		for count > 0 {
			node := tempStore[0]
			tempStore = tempStore[1:]
			tempRes = append(tempRes, node.Val)
			if node.Left != nil {
				tempStore = append(tempStore, node.Left)
			}
			if node.Right != nil {
				tempStore = append(tempStore, node.Right)
			}
			count = count - 1
		}
		res = append(res, tempRes)
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

	res := levelOrder(a)
	fmt.Println(res)

}
