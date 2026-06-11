package LeetCode

import (
	"slices"
	"sort"
)

// 二分
func smallestDistancePair(nums []int, k int) int {
	slices.Sort(nums)
	n := len(nums)
	check := func(mx int) bool {
		count := 0
		for i := 0; i < n; i++ {
			index := sort.SearchInts(nums, nums[i]+mx+1) - 1
			count += index - i
		}
		return count >= k
	}
	l := 0
	r := slices.Max(nums) - slices.Min(nums)
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

// 滑动窗口 + 二分
func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	n := len(nums)
	check := func(mx int) bool {
		count := 0
		l := 0
		//滑动窗口求子序列 越短越合法
		for r := 1; r < n; r++ {
			for nums[r]-nums[l] > mx {
				l++
			}
			count += r - l
		}
		return count >= k
	}
	l := 0
	r := slices.Max(nums) - slices.Min(nums)
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
