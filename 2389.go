package LeetCode

import "slices"

func answerQueries(nums []int, queries []int) []int {
	//排序后前缀和 + 贪心 (要想和尽可能小那就需要选最小的几个数加进去)
	slices.Sort(nums)
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	//二分 查找小于等于target的最大值
	for i := 0; i < len(queries); i++ {
		queries[i] = findTarget(nums, queries[i])
	}
	return queries
}
func findTarget(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return r
}
