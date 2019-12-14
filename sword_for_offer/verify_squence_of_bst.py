'''
输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历的结果。如果是则输出Yes,否则输出No。假设输入的数组的任意两个数字都互不相同。
'''

class Solution:
    def VerifySquenceOfBST(self, sequence):
        if len(sequence) == 0 :
            return False
        return self.Recursion(sequence)


    def Recursion(self, sequence):
        if len(sequence) == 1 or len(sequence) == 0:
            return True
        head_num = sequence[-1]
        right_head = len(sequence) - 1
        for i in range(0,len(sequence)-1):
            if sequence[i] > head_num:
                right_head = i
                break

        for i in range(right_head,len(sequence)-1):
            if head_num > sequence[i]:
                return False

        left = self.Recursion(sequence[0:right_head])
        right = self.Recursion(sequence[right_head:len(sequence)-1])

        return left and right