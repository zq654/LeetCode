package LeetCode

import "sort"

// todo 解法1 自己写的dp 比较慢 。解法二 贪心 二分 DP比较快
func lengthOfLIS(nums []int) int {
	length := len(nums)
	if length == 1 {
		return 1
	}
	dp := make([]int, length)
	dp[0] = 1
	res := 0
	for i := 1; i < length; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			} else {
				dp[i] = max(dp[i], 1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}

func lengthOfLIS(nums []int) int {
	g := nums[:0] // 原地修改
	for _, x := range nums {
		j := sort.SearchInts(g, x)
		if j == len(g) { // >=x 的 g[j] 不存在
			g = append(g, x)
		} else {
			g[j] = x
		}
	}
	return len(g)
}
