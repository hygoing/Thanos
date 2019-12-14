# Definition for a Node.
class Node(object):
    def __init__(self, val, next, random):
        self.val = val
        self.next = next
        self.random = random


class Solution(object):
    def copyRandomList(self, head):
        if head is None:
            return None

        node = head
        while node is not None:
            new_node = Node(node.val, node.next, None)
            node.next = new_node
            node = new_node.next

        node = head
        while node is not None:
            new_node = node.next
            if node.random is not None:
                new_node.random = node.random.next
            else:
                new_node.random = None
            node = new_node.next

        node = head
        new_head = head.next
        while node is not None:
            new_node = node.next
            node.next = new_node.next
            if new_node.next is not None:
                new_node.next = new_node.next.next
            else:
                new_node.next = None
            node = node.next

        return new_head
