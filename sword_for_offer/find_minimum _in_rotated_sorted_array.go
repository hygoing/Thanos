package main

func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := (left + right) / 2
		if nums[mid] < nums[right] {
			right = mid
		} else if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = right - 1
		}
	}
	return nums[left]
}

func findMin1(nums []int) int {
	left := 0
	right := len(nums) - 1
	mid := 0

	for nums[left] >= nums[right] {
		mid = (left + right) >> 1

		if right-left == 1 {
			return nums[right]
		}

		if nums[left] == nums[right] && nums[left] == nums[mid] {
			min := nums[left]
			for i := left; i <= right; i++ {
				if min > nums[i] {
					min = nums[i]
				}
			}
			return min
		}

		if nums[mid] >= nums[right] {
			left = mid
		} else {
			right = mid
		}
	}
	return nums[mid]
}
