package LeetCode

import "fmt"

func findMaxLength(nums []int) int {
	recordMap := make(map[int]int, len(nums)) //记录0与1的差值出现的下标
	recordMap[0] = -1
	res := 0
	sum := 0 //0的个数减1的个数
	for i, num := range nums {
		sum += 1 - 2*num
		if index, ok := recordMap[sum]; ok {
			fmt.Println(index)
			res = max(res, i-index)
		} else {
			recordMap[sum] = i
		}
	}
	return res
}
