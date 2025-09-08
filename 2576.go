package LeetCode

import (
	"slices"
)

// todo 贪心没做出来
func maxNumOfMarkedIndices(nums []int) int {
	slices.Sort(nums)
	l := 0
	r := (len(nums) + 1) / 2
	res := 0
	for l <= (len(nums)+1)/2 && r < len(nums) {
		if nums[r] >= nums[l]*2 {
			res += 2
			r++
			l++
		} else {
			r++
		}
	}
	return res
}
