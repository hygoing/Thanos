package main

import "sort"

/*
给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。

例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]
 */


func threeSum(nums []int) [][]int {
	answers := make([][]int, 0)
	if len(nums) < 3 {
		return answers
	}

	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		L := i + 1
		R := len(nums) - 1
		for L < R {
			sum := nums[i] + nums[L] + nums[R]
			if sum == 0 {
				answers = append(answers, []int{nums[i], nums[L], nums[R]})
				for L < R && nums[L] == nums[L+1] {
					L++
				}
				for R > L && nums[R] == nums[R-1] {
					R--
				}
				L++
				R--
			}
			if sum > 0 {
				R--
			}
			if sum < 0 {
				L++
			}
		}
	}
	return answers
}
