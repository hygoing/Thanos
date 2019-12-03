package main

import . "persion.com/thanos/common"

func buildTreeByPreIn(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 || len(inorder) != len(preorder) {
		return nil
	}
	return buildTreeByPreInRecursion(preorder, inorder)
}

func buildTreeByPreInRecursion(preoder []int, inorder []int) *TreeNode {
	head := preoder[0]
	head_inorder_index := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == head {
			head_inorder_index = i
		}
	}

	left_inorder := inorder[:head_inorder_index]
	right_inorder := inorder[head_inorder_index+1:]
	left_preorder := preoder[1:len(left_inorder)+1]
	right_preorder := preoder[len(left_inorder)+1:]

	node := &TreeNode{
		Val: head,
	}

	if len(left_inorder) == 0 && len(right_inorder) == 0 {
		return node
	}

	if len(left_inorder) > 0 {
		node.Left = buildTreeByPreInRecursion(left_preorder, left_inorder)
	}
	if len(right_inorder) > 0 {
		node.Right = buildTreeByPreInRecursion(right_preorder, right_inorder)
	}
	return node
}
