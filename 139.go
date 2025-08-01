package LeetCode

func wordBreak(s string, wordDict []string) bool {
	//空间换时间用map记录wordDict
	wordDictMap := make(map[string]bool)
	for _, word := range wordDict {
		wordDictMap[word] = true
	}
	length := len(s)
	bytes := make([]byte, length+1)
	dp := make([]bool, length+1)
	for i := 1; i < length+1; i++ {
		bytes[i] = s[i-1]
	}
	dp[0] = true
	for i := 1; i < length+1; i++ {
		for j := i - 1; j >= 0; j-- {
			dp[i] = dp[j] && wordDictMap[s[j:i]]
			if dp[i] {
				break
			}
		}
	}
	return dp[length]
}
