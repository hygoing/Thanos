package main

import . "persion.com/thanos/common"

func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	tempStore := make([]*TreeNode, 0)

	if root == nil {
		return res
	}

	tempStore = append(tempStore, root)
	for len(tempStore) > 0 {
		node := tempStore[len(tempStore) - 1]
		tempStore = tempStore[:len(tempStore) - 1]
		res = append([]int{node.Val}, res[:]...)

		if node.Left != nil {
			tempStore = append(tempStore, node.Left)
		}
		if node.Right != nil {
			tempStore = append(tempStore, node.Right)
		}
	}
	return res
}