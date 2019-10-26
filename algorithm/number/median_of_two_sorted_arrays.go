package main

import (
	"fmt"
	"math"
)

/*
给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。

请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

你可以假设 nums1 和 nums2 不会同时为空。

示例 1:

nums1 = [1, 3]
nums2 = [2]

则中位数是 2.0
示例 2:

nums1 = [1, 2]
nums2 = [3, 4]

则中位数是 (2 + 3)/2 = 2.5
 */

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	left := (m + n + 1) / 2
	right := (m + n + 2) / 2

	return (getKth(nums1, 0, m-1, nums2, 0, n-1, left) + getKth(nums1, 0, m-1, nums2, 0, n-1, right)) / 2
}

func getKth(nums1 []int, begin_1 int, end_1 int, nums2 []int, begin_2 int, end_2 int, k int) float64 {
	len_1 := end_1 - begin_1 + 1
	len_2 := end_2 - begin_2 + 1

	if len_1 > len_2 {
		return getKth(nums2, begin_2, end_2, nums1, begin_1, end_1, k)
	}
	if len_1 == 0 {
		return float64(nums2[begin_2+k-1])
	}
	if k == 1 {
		return math.Min(float64(nums1[begin_1]), float64(nums2[begin_2]))
	}

	i := begin_1 + int(math.Min(float64(len_1), float64(k/2))) - 1
	j := begin_2 + int(math.Min(float64(len_2), float64(k/2))) - 1

	if nums1[i] < nums2[j] {
		return getKth(nums1, i+1, end_1, nums2, begin_2, end_2, k-(i-begin_1+1))
	} else {
		return getKth(nums1, begin_1, end_1, nums2, j+1, end_2, k-(j-begin_2+1))
	}
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	answer := findMedianSortedArrays(nums1, nums2)
	fmt.Println(answer)

}

/*
answer 2

	m := len(nums1)
	n := len(nums2)
	left, right := -1, -1
	i, j := 0, 0

	for k := 0; k <= (m+n)/2; k++ {
		left = right
		if i < m && (j >= n || nums1[i] < nums2[j]) {
			right = nums1[i]
			i++
		} else {
			right = nums2[j]
			j++
		}
	}

	if (m+n )%2 == 0 {
		return (float64(left) + float64(right)) / 2
	} else {
		return float64(right)
	}

 */

/*
answer 1

answer := 0
	i := 0
	j := 0
	length := len(nums1) + len(nums2)
	counter := 0
	isEVen := false
	if (length)%2 == 0 {
		isEVen = true
	}

	for i < len(nums1) && j < len(nums2) {
		num := 0
		counter++
		if nums1[i] < nums2[j] {
			num = nums1[i]
			i++
		} else {
			num = nums2[j]
			j++
		}

		if isEVen && counter == length/2 {
			answer = answer + num
		}
		if counter == length/2+1 {
			answer = answer + num
			break
		}
	}


	for i < len(nums1) {
		counter++
		num :=nums1[i]
		i++
		if isEVen && counter == length/2 {
			answer = answer + num
		}
		if counter == length/2+1 {
			answer = answer + num
			break
		}
	}

	for j < len(nums2) {
		counter++
		num :=nums2[j]
		j++
		if isEVen && counter == length/2 {
			answer = answer + num
		}
		if counter == length/2+1 {
			answer = answer + num
			break
		}
	}

	fmt.Println(answer)

	if isEVen {
		return float64(answer) / float64(2)
	} else {
		return float64(answer)
	}

 */
