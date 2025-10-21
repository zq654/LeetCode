package LeetCode

import "slices"

func minimizeMax(nums []int, p int) int {
	if p == 0 {
		return 0
	}
	slices.Sort(nums)
	length := len(nums)
	check := func(mx int) bool {
		count := 0
		for i := 1; i < length; i++ {
			if nums[i]-nums[i-1] <= mx {
				count++
				i++
			}
			if count >= p {
				return true
			}
		}
		return false
	}
	l := 0
	r := nums[length-1] - nums[0]
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
