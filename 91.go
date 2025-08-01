package LeetCode

// todo 这题纯让我试出来的 还没想清楚为什么是 dp[i-1]+dp[i-2]
func numDecodings(s string) int {
	//dp[i] 代表当前选到当前时的数量
	bytes := []byte(s)
	dp := make([]int, len(bytes))
	dp[0] = 1
	if bytes[0] == '0' {
		return 0
	}
	for i := 1; i < len(dp); i++ {
		count := (bytes[i-1]-'0')*10 + (bytes[i] - '0')
		if bytes[i] == '0' && (count < 10 || count > 26) {
			return 0
		}
		if bytes[i] == '0' && count >= 10 && count <= 26 {
			if i-2 < 0 {
				dp[i] = 1
			} else {
				dp[i] = dp[i-2]
			}
		}
		if bytes[i] != '0' && count >= 10 && count <= 26 {
			if i-2 < 0 {
				dp[i] = dp[i-1] + 1
			} else {
				dp[i] = dp[i-1] + dp[i-2]
			}
		}
		if bytes[i] != '0' && (count < 10 || count > 26) {
			dp[i] = dp[i-1]
		}
	}
	return dp[len(dp)-1]
}
