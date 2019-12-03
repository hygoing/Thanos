package main

import . "persion.com/thanos/common"

func constructFromPrePost(pre []int, post []int) *TreeNode {
	if len(pre) == 0 || len(post) == 0 || len(pre) != len(post) {
		return nil
	}
	return buildTreeByPrePostRecursion(pre, post)
}

func buildTreeByPrePostRecursion(pre []int, post []int) *TreeNode {
	node := &TreeNode{
		Val: pre[0],
	}

	if len(pre) == 1 && len(post) == 1 {
		return node
	}

	head_right := post[len(post) - 2]
	head_right_pre_index := 0
	for i := 0; i < len(pre); i ++ {
		if pre[i] == head_right {
			head_right_pre_index = i
		}
	}

	left_pre := pre[1:head_right_pre_index]
	right_pre := pre[head_right_pre_index:]
	left_post := post[:len(left_pre)]
	right_post := post[len(left_pre):len(post) - 1]

	if len(left_pre) > 0 {
		node.Left = buildTreeByPrePostRecursion(left_pre, left_post)
	}
	if len(right_pre) > 0 {
		node.Right = buildTreeByPrePostRecursion(right_pre, right_post)
	}
	return node
}