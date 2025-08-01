package LeetCode

// 自己写的边界花了不少时间
func findTargetSumWays(nums []int, target int) int {
	//dp[i][j] 表示选到第i个值的时候 值为j的 可能数量
	//因为target可能有复数 所以先计算count
	count := 0
	length := len(nums)
	for i := 0; i < length; i++ {
		count = count + nums[i]
	}
	if count < target || -count > target {
		return 0
	}

	//计算值 可能范围是[-count,count]
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, 2*count+1)
	}
	//所有值加一个count
	dp[0][count+nums[0]] = 1
	dp[0][count-nums[0]] = 1
	if nums[0] == 0 {
		dp[0][count] = 2
	}

	for i := 1; i < length; i++ {
		for j := 0; j < 2*count+1; j++ {
			if j < nums[i] {
				dp[i][j] = dp[i-1][j+nums[i]]
			} else if j >= 2*count+1-nums[i] {
				dp[i][j] = dp[i-1][j-nums[i]]
			} else {
				dp[i][j] = dp[i-1][j-nums[i]] + dp[i-1][j+nums[i]]
			}
		}
	}
	return dp[length-1][target+count]
}
