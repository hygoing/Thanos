package main

import . "persion.com/thanos/common"

/*
翻转一棵二叉树。

示例：

输入：

     4
   /   \
  2     7
 / \   / \
1   3 6   9
输出：

     4
   /   \
  7     2
 / \   / \
9   6 3   1
*/

func invertTree(root *TreeNode) *TreeNode {
	invertRecursion(root)
	return root
}

func invertRecursion(node *TreeNode) {
	if node == nil {
		return
	}
	temp := node.Left
	node.Left = node.Right
	node.Right = temp
	invertRecursion(node.Left)
	invertRecursion(node.Right)
}
