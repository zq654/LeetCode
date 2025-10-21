package LeetCode

import "slices"

func minimizeArrayValue(nums []int) int {
	//从0 一直到len(nums)-2 往后面索取
	check := func(mx int) bool {
		count := 0
		for i := 0; i <= len(nums)-2; i++ {
			if nums[i]-count > mx {
				return false
			}
			//当前可以再容纳的值为
			count = mx - nums[i] + count
		}
		return nums[len(nums)-1]-count <= mx
	}
	l := slices.Min(nums)
	r := slices.Max(nums)
	for l <= r {
		mid := l + (r-l)>>1
		if check(mid) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}
