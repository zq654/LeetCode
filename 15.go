package LeetCode

import (
	"slices"
)

func threeSum(nums []int) [][]int {
	//求nums[i] + nums[j] + nums[k] == 0 ==> 求num[j] + nums[k] == -nums[i]
	res := make([][]int, 0)
	set := make(map[[3]int]bool)
	slices.Sort(nums)
	for i := 0; i < len(nums)-3; i++ {
		j := i + 1
		k := len(nums) - 1
		for j < k {

			if nums[j]+nums[k] > -nums[i] {
				k--
			} else if nums[j]+nums[k] < -nums[i] {
				j++
			} else {
				if !set[[3]int{nums[i], nums[j], nums[k]}] {
					res = append(res, []int{nums[i], nums[j], nums[k]})
					set[[3]int{nums[i], nums[j], nums[k]}] = true
				}
				j++
				k--
			}
		}
	}
	return res
}
