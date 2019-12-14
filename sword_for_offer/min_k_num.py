class Solution:
    def GetLeastNumbers_Solution(self, tinput, k):
        stack = []
        count = 0
        for i in range(0, len(tinput)):
            if len(stack) == 0:
                stack.append(tinput[i])
                count = count + 1
            else:
                temp = []
                if count < k:
                    while len(stack) > 0 and stack[-1] > tinput[i]:
                        temp.append(stack[-1])
                        stack = stack[:len(stack)-1]
                    stack.append(tinput[i])
                    for j in range(len(temp) - 1, -1, -1):
                        stack.append(temp[j])
                    count = count + 1
                elif stack[-1] > tinput[i]:
                    while len(stack) > 0 and stack[-1] > tinput[i]:
                        temp.append(stack[-1])
                        stack = stack[:len(stack)-1]
                    stack = stack[1:len(stack)]
                    stack.append(tinput[i])
                    for j in range(len(temp) - 1, -1, -1):
                        stack.append(temp[j])
        return stack


if __name__ == "__main__":
    s = Solution()
    nums = [4,5,1,6,2,7,3,8]
    res = s.GetLeastNumbers_Solution(nums,4)
    print(res)
