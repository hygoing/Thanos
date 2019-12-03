package main

func fibIteration(N int) int {
	if N < 2 {
		return N
	}
	x, y, res := 0, 1, 0
	for i := 2; i <= N; i++ {
		res = x + y
		x = y
		y = res
	}
	return res
}

func fibIteration2(N int) int {
	pre, next := 0, 1
	for N > 0 {
		next = next + pre
		pre = next - pre
		N--
	}
	return pre
}

func fibRecursion(N int) int {
	if N < 2 {
		return N
	}
	return fibRecursion(N-1) + fibRecursion(N-2)
}

func fibDp(N int) int {
	if N < 2 {
		return N
	}

	tempStore := []int{0, 1}
	for i := 2; i <= N; i++ {
		tempStore = append(tempStore, tempStore[i-1]+tempStore[i-2])
	}
	return tempStore[N]
}
