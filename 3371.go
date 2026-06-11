package LeetCode

import "math"

func getLargestOutlier(nums []int) int {
	//拆解题意 一个数可以是其他数的和 剩余的哪个值就是我们要的target
	//可得 (sum-target)/2 的值存在（且不能是自己）
	sum := 0
	recordMap := make(map[int]int)
	for _, num := range nums {
		sum += num
		recordMap[num]++
	}
	res := math.MinInt64
	for k, _ := range recordMap {

		if (sum-k)%2 == 0 && (recordMap[(sum-k)/2] > 1 || (recordMap[(sum-k)/2] > 0 && 3*k != sum)) {
			res = max(res, k)
		}
	}
	return res
}
