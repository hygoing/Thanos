class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def Convert(self, pRootOfTree):
        head, node = pRootOfTree, pRootOfTree
        while pRootOfTree is not None:
            head = pRootOfTree
            pRootOfTree = pRootOfTree.left

        tempStack = []
        while node is not None or len(tempStack) > 0:
            while node is not None:
                tempStack.append(node)
                node = node.left

            node = tempStack[-1]
            tempStack = tempStack[:len(tempStack) - 1]
            if node.left is None and node.right is None:
                if len(tempStack) > 0 and tempStack[-1].val > node.val:
                    node.right = tempStack[-1]
                    tempStack[-1].left = node
                node = None
            else:
                temp = node.right
                while temp is not None:
                    tempStack.append(temp)
                    temp = temp.left
                if len(tempStack) > 0:
                    node.right = tempStack[-1]
                    tempStack[-1].left = node

                node = temp
        return head


if __name__ == "__main__":
    node_10 = TreeNode(10)
    node_6 = TreeNode(6)
    node_14 = TreeNode(14)
    node_4 = TreeNode(4)
    node_8 = TreeNode(8)
    node_12 = TreeNode(12)
    node_16 = TreeNode(16)

    node_10.left = node_4
    node_10.right = node_14
    node_6.left = node_4
    node_6.right = node_8
    node_14.left = node_12
    node_14.right = node_16

    s = Solution()
    head = s.Convert(node_10)
