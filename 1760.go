package LeetCode

import "slices"

func minimumSize(nums []int, maxOperations int) int {
	check := func(mx int) bool {
		count := 0
		for i := 0; i < len(nums); i++ {
			count += ((nums[i] + mx - 1) / mx) - 1
			if count > maxOperations {
				return false
			}
		}
		return true
	}
	l := 1
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
