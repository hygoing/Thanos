package main

func hammingWeightUint(num uint32) int {
	res := 0
	for num != 0 {
		if num&1 == 1 {
			res = res + 1
		}
		num = num >> 1
	}
	return res
}

func hammingWeightInt(num int) int {
	res := 0
	flag := 1
	for flag != 0 {
		if num&flag == 1 {
			res = res + 1
		}
		flag = flag << 1
	}
	return res
}

func hammingWeightIntNb(num int) int {
	res := 0
	for num != 0 {
		res = res + 1
		num = (num - 1) & num
	}
	return res
}
