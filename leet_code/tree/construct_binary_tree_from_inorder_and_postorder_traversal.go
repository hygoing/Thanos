package main

import . "persion.com/thanos/common"

func buildTreeByInPost(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 || len(inorder) != len(postorder) {
		return nil
	}

	return buildTreeByInPostRecursion(inorder, postorder)

}

func buildTreeByInPostRecursion(inorder []int, postorder []int) *TreeNode {
	head := postorder[len(postorder) - 1]
	head_inorder_index := 0
	for i:=0; i < len(inorder); i++ {
		if inorder[i] == head {
			head_inorder_index = i
		}
	}

	left_inorder := inorder[:head_inorder_index]
	right_inorder := inorder[head_inorder_index + 1:]
	left_postorder := postorder[:len(left_inorder)]
	right_postorder := postorder[len(left_inorder):len(postorder) - 1]

	node := &TreeNode{
		Val: head,
	}

	if len(left_inorder) == 0 && len(right_inorder) == 0 {
		return node
	}

	if len(left_inorder) > 0 {
		node.Left = buildTreeByInPostRecursion(left_inorder, left_postorder)
	}

	if len(right_inorder) > 0 {
		node.Right = buildTreeByInPostRecursion(right_inorder, right_postorder)
	}
	return node
}