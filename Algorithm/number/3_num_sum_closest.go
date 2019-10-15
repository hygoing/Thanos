package main

/*
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。

例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.

与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).
 */

import (
	"sort"
	"math"
)

func threeSumClosest(nums []int, target int) int {
	answer := 0
	if len(nums) < 3 {
		return answer
	}
	answer = nums[0] + nums[1] + nums[2]
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		L := i + 1
		R := len(nums) - 1
		for L < R {
			sum := nums[i] + nums[L] + nums[R]
			if int(math.Abs(float64(sum-target))) <= int(math.Abs(float64(target-answer))) {
				answer = sum
			}
			if sum == target {
				return answer
			}
			if sum > target {
				R--
			}
			if sum < target {
				L++
			}
		}
	}
	return answer
}
