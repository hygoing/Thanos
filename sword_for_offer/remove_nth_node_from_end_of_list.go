package main

import . "persion.com/thanos/common"

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n == 0 {
		return nil
	}

	guard := &ListNode{
		Next: head,
	}
	prev := guard
	curr := guard
	for i := 0; i < n; i++ {
		if curr.Next != nil {
			curr = curr.Next
		}else {
			return nil
		}
	}

	for curr.Next != nil {
		prev = prev.Next
		curr = curr.Next
	}
	prev.Next = prev.Next.Next

	return guard.Next
}
