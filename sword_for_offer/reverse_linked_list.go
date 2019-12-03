package main

import "persion.com/thanos/common"

/*
反转一个单链表。

示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
*/

func reverseListIteration(head *common.ListNode) *common.ListNode {
	if head == nil {
		return nil
	}
	var HEAD *common.ListNode = nil

	for head != nil {
		next := head.Next
		head.Next = HEAD
		HEAD = head
		head = next
	}
	return HEAD
}

func reverseList(head *common.ListNode) *common.ListNode {
	if head == nil {
		return nil
	}
	return recursion(head)
}

func recursion(node *common.ListNode) *common.ListNode {
	if node.Next == nil {
		return node
	}
	head := recursion(node.Next)
	node.Next.Next = node
	node.Next = nil
	return head
}
