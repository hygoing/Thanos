package main

import "fmt"

func binaryChopIteration(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}

	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		num := nums[mid]
		if num == target {
			return true
		} else if num > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return false
}

func binaryChopRecursion(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	left := 0
	right := len(nums) - 1
	mid := (left + right) / 2
	num := nums[mid]
	if num == target {
		return true
	} else if num > target {
		return binaryChopIteration(nums[:mid], target)
	} else {
		return binaryChopIteration(nums[mid+1:], target)
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	res := binaryChopIteration(nums, 10)
	fmt.Println(res)

	res2 := binaryChopRecursion(nums, 0)
	fmt.Println(res2)
}
