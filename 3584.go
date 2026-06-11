package LeetCode

import "math"

func maximumProduct(nums []int, m int) int64 {
	//删除一些值之后子序列长度为m 可得求满足序列长度大于等于m的边界最大乘积
	r := m - 1
	l := 0
	maxValue := nums[0]
	minValue := nums[0]
	res := math.MinInt64
	for ; r < len(nums); r++ {
		maxValue = max(maxValue, nums[l])
		minValue = min(minValue, nums[l])
		res = max(nums[r]*maxValue, nums[r]*minValue, res)
		l++
	}
	return int64(res)
}
