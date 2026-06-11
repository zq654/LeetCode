package LeetCode

import "math"

func maximumDifference(nums []int) int {

	minValue := math.MaxInt64
	res := -1
	for _, num := range nums {
		res = max(res, num-minValue)
		if num < minValue {
			minValue = num
		}
	}
	if res == 0 {
		return -1
	}
	return res
}
