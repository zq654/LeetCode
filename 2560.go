package LeetCode

import "slices"

func minCapability(nums []int, k int) int {
	//dp + 二分
	length := len(nums)
	if length <= 2 {
		return slices.Min(nums)
	}
	check := func(mx int) bool {
		dp := make([]int, length+1) //dp[i]表示到第i家最多能偷多少个
		dp[0] = 0
		if nums[0] <= mx {
			dp[1] = 1
		} else {
			dp[1] = 0
		}
		if nums[1] <= mx {
			dp[2] = 1
		} else {
			dp[2] = 0
		}
		for i := 3; i <= length; i++ {
			if nums[i-1] <= mx {
				dp[i] = max(dp[i-3]+1, dp[i-2]+1, dp[i-1])
			} else {
				dp[i] = max(dp[i-3], dp[i-2], dp[i-1])
			}
			if dp[i] >= k {
				return true
			}
		}
		return false
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
