package main

/*
给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

返回滑动窗口中的最大值。

 

示例:

输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
输出: [3,3,5,5,6,7]
解释:

  滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
 */

func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, 0)
	temp := make([]int, 0)
	if len(nums) < k {
		return res
	}

	for i := 0; i < len(nums); i++ {
		for len(temp) > 0 && nums[temp[len(temp)-1]] <= nums[i] {
			temp = temp[:len(temp)-1]
		}
		temp = append(temp, i)

		if temp[0] == i-k {
			temp = temp[1:]
		}


		// 当数组大于等于窗口时，每次循环都需要产生一个结果
		if i >= k-1 {
			res = append(res, nums[temp[0]])
		}

	}
	return res
}
