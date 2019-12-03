package main

func sortArrayByParity(A []int) []int {
	left := 0
	right := len(A) - 1

	for left < right {
		if A[left]&1 == 1 && A[right]&1 == 0 {
			temp := A[left]
			A[left] = A[right]
			A[right] = temp
		}
		if A[left]&1 == 0 {
			left++
		}
		if A[right]&1 == 1 {
			right--
		}
	}
	return A
}

func sortArrayByParity1(A []int) []int {
	left := 0
	right := len(A) - 1

	for left < right {
		for left < right && A[left]&1 == 0 {
			left++
		}
		for left < right && A[right]&1 == 1 {
			right--
		}
		if left < right {
			temp := A[left]
			A[left] = A[right]
			A[right] = temp
		}
	}
	return A
}