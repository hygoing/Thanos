/*
前几个月放映的头号玩家简直火得不能再火了，作为一个探索终极AI的研究人员，月神自然去看了此神剧。
由于太过兴奋，晚上月神做了一个奇怪的梦，月神梦见自己掉入了一个被施放了魔法的深渊，月神想要爬上此深渊。

已知深渊有N层台阶构成（1 <= N <= 1000)，并且每次月神仅可往上爬2的整数次幂个台阶(1、2、4、....)，请你编程告诉月神，月神有多少种方法爬出深渊
输入描述:
输入共有M行，(1<=M<=1000)

第一行输入一个数M表示有多少组测试数据，

接着有M行，每一行都输入一个N表示深渊的台阶数
输出描述:
输出可能的爬出深渊的方式
 */

package main

import (
	"fmt"
)

func main() {
	var dataCount = 0
	fmt.Scanln(&dataCount)

	var stepLengths []int
	var stepLength int
	for i := 0; i < dataCount; i++ {
		fmt.Scanln(&stepLength)
		stepLengths = append(stepLengths, stepLength)
	}

	//用来存储结果
	var result = make([]int, 0)
	//用来存储数组下标台阶数时爬出深渊的方式数目
	stepToCount := []int{1, 1, 2}
	for _, stepLength := range stepLengths {
		if stepLength == 1 {
			result = append(result, 1)
			continue
		}
		if stepLength == 2 {
			result = append(result, 2)
			continue
		}
		for tempStepLength := len(stepToCount); tempStepLength <= stepLength; tempStepLength++ {
			tempStepCount := 0
			for base := 1; base <= tempStepLength; base = base * 2 {
				tempStepCount = tempStepCount + stepToCount[tempStepLength-base]
			}
			if len(stepToCount) == tempStepLength {
				stepToCount = append(stepToCount, tempStepCount % 1000000003)
			}
		}
		result = append(result, stepToCount[stepLength])
	}

	for _, res := range result {
		fmt.Println(res)
	}
}
