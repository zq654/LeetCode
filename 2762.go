package LeetCode

import "math"

// todo 思路是对的 但是没有实现
func continuousSubarrays(nums []int) int64 {
	maxValue := -1
	minValue := math.MaxInt //记录窗口内的最大值和最小值 只要最大值和最小值之间的差小于等于2 则其他值的差肯定也小于等于2
	//用数组记录
	l := 0
	res := 0
	for r := 0; r < len(nums); r++ {
		maxValue = max(maxValue, nums[r])
		minValue = min(minValue, nums[r])
		for maxValue-minValue > 2 {

		}
	}
}
