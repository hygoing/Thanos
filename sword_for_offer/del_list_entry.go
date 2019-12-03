package main

import . "persion.com/thanos/common"

func removeElements1(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}
	if head == nil {
		return head
	}
	current := head
	for current.Next != nil {
		if current.Next.Val == val {
			current.Next  = current.Next.Next
			continue
		}
		current = current.Next
	}
	return head
}


func removeElements2(head *ListNode, val int) *ListNode {
	guard := &ListNode{
		Next:head,
	}
	current := guard
	for current.Next != nil {
		if current.Next.Val == val {
			current.Next = current.Next.Next
			continue
		}
		current = current.Next
	}
	return guard.Next
}
