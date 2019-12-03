package main

import . "persion.com/thanos/common"

func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	recursion(&res, root)
	return res
}

func recursion(res *[]int, node *TreeNode) {
	if node == nil {
		return
	}
	*res = append(*res, node.Val)
	recursion(res, node.Left)
	recursion(res, node.Right)
}

func preorderTraversalIteration(root *TreeNode) []int {
	result := make([]int, 0)
	tempStore := make([]*TreeNode, 0)

	if root == nil {
		return result
	}

	tempStore = append(tempStore, root)
	for len(tempStore) > 0 {
		node := tempStore[len(tempStore)-1]
		result = append(result, node.Val)
		tempStore = tempStore[:len(tempStore)-1]
		if node.Right != nil {
			tempStore = append(tempStore, node.Right)
		}
		if node.Left != nil {
			tempStore = append(tempStore, node.Left)
		}
	}

	return result
}

func preorderTraversalMorris(root *TreeNode) []int {
	res := make([]int, 0)

	if root == nil {
		return res
	}

	node := root
	for node != nil {
		if node.Left == nil {
			res = append(res, node.Val)
			node = node.Right
		} else {
			predecessor := node.Left
			for predecessor.Right != nil && predecessor.Right != node {
				predecessor = predecessor.Right
			}

			if predecessor.Right == nil {
				res = append(res, node.Val)
				predecessor.Right = node
				node = node.Left
			} else {
				predecessor.Right = nil
				node = node.Right
			}
		}
	}
	return res
}
