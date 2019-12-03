package main

import "fmt"

func power(x float64, n int) float64 {
	absN := n
	if n < 0 {
		x = 1 / x
		absN = -n
	}
	return powerRecursion(x, absN)
}

func powerRecursion(base float64, exponent int) float64 {
	if exponent == 0 {
		return 1
	}
	if exponent == 1 {
		return base
	}
	result := powerRecursion(base, exponent>>1)
	result = result * result
	if exponent&1 == 1 {
		result = result * base
	}
	return result
}

func powerIteration(x float64, n int) float64 {
	absN := n
	if n < 0 {
		x = 1 / x
		absN = -n
	}
	result := 1.0
	tempResult := x
	for i := absN; i > 0; i = i >> 1 {
		if absN&1 == 1 {
			result = result * tempResult
		}
		tempResult = tempResult * tempResult
	}
	return result
}

func main() {
	fmt.Println(power(2.0, 6))
}
