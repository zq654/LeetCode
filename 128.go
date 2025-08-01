package LeetCode

import (
	"math"
)

func LongestConsecutive(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	m := make(map[int]struct{})
	minNums := math.MaxInt32
	maxNums := math.MinInt32
	res := 1
	length := 1
	stepLength := 0
	for _, num := range nums {
		m[num] = struct{}{}
		if num > maxNums {
			maxNums = num
		}
		if num < minNums {
			minNums = num
		}
	}
	for i := minNums + 1; i <= maxNums; i++ {
		if _, ok := m[i]; ok {
			length = length + 1
		} else {
			stepLength++
			res = max(res, length)
			length = 0
		}
		if stepLength >= len(nums) {
			break
		}
	}
	res = max(res, length)
	return res
}
