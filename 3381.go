package LeetCode

import "math"

func maxSubarraySum(nums []int, k int) int64 {
	// 前缀和 + dp
	length := len(nums)
	dp := make([]int, length+1)
	pre := make([]int, length+1)
	dp[0] = 0
	pre[0] = 0
	sum := 0
	res := math.MinInt64
	for i, num := range nums {
		sum += num
		pre[i+1] = sum
		if i >= k-1 {
			value := pre[i+1] - pre[i+1-k]
			dp[i+1] = max(dp[i+1-k]+value, value)
			res = max(res, dp[i+1])
		}
	}
	return int64(res)
}

// 灵神枚举+前缀和
func maxSubarraySum(nums []int, k int) int64 {
	minS := make([]int, k)
	for i := 0; i < k-1; i++ {
		minS[i] = math.MaxInt / 2 // 防止下面减法溢出
	}

	ans := math.MinInt
	s := 0
	for j, x := range nums {
		s += x
		i := j % k
		ans = max(ans, s-minS[i])
		minS[i] = min(minS[i], s)
	}
	return int64(ans)
}
