package main

/*
给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。

说明:

初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
示例:

输入:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

输出: [1,2,2,3,5,6]
 */

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	ptr_1 := m - 1
	ptr_2 := n - 1
	ptr := m + n - 1

	for ptr_1 >= 0 && ptr_2 >= 0 {
		if nums1[ptr_1] > nums2[ptr_2] {
			nums1[ptr] = nums1[ptr_1]
			ptr_1--
		} else {
			nums1[ptr] = nums2[ptr_2]
			ptr_2--
		}
		ptr--
	}
	for i := 0; i <= ptr_2; i++ {
		nums1[i] = nums2[i]
	}
	fmt.Println(nums1)
}


func merge1(nums1 []int, m int, nums2 []int, n int)  {
	ptr_1 := m - 1
	ptr_2 := n -1
	ptr := m + n - 1

	for ptr >=0 {
		if ptr_1 >= 0 && (ptr_2 < 0 || nums1[ptr_1] > nums2[ptr_2]){
			nums1[ptr] = nums1[ptr_1]
			ptr_1 = ptr_1 - 1
		}else{
			nums1[ptr] = nums2[ptr_2]
			ptr_2 = ptr_2 - 1
		}
		ptr--
	}
	fmt.Print(nums1)
}