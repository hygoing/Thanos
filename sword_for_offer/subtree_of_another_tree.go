package main

import . "persion.com/thanos/common"

/*
给定两个非空二叉树 s 和 t，检验 s 中是否包含和 t 具有相同结构和节点值的子树。s 的一个子树包括 s 的一个节点和这个节点的所有子孙。s 也可以看做它自身的一棵子树。

示例 1:
给定的树 s:

     3
    / \
   4   5
  / \
 1   2
给定的树 t：

   4
  / \
 1   2
返回 true，因为 t 与 s 的一个子树拥有相同的结构和节点值。
 */

func isSubtree(s *TreeNode, t *TreeNode) bool {
	res := false
	if s != nil && t != nil {
		if s.Val == t.Val {
			res = subRecursion(s, t)
		}
		if res == false {
			res = isSubtree(s.Left, t)
		}
		if res == false {
			res = isSubtree(s.Right, t)
		}
	}
	return res
}

func subRecursion(s *TreeNode, t *TreeNode) bool {
	if t == nil && s == nil {
		return true
	}
	if t == nil || s == nil {
		return false
	}
	if t.Val != s.Val {
		return false
	}
	return subRecursion(s.Left, t.Left) && subRecursion(s.Right, t.Right)
}

func fetchInorder(t *TreeNode, inorder *[]int) {
	if t == nil {
		return
	}
	fetchInorder(t.Left, inorder)
	*inorder = append(*inorder, t.Val)
	fetchInorder(t.Right, inorder)
}
