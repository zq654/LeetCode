package LeetCode

func lengthOfLongestSubstring(s string) int {
	recordMap := make(map[byte]int) //记录该字符出现的上一个位置
	//滑动窗口左边界
	l := 0
	res := 0
	for i := 0; i < len(s); i++ {
		if index, ok := recordMap[s[i]]; index >= l && ok {

			l = recordMap[s[i]] + 1
		}
		recordMap[s[i]] = i
		res = max(res, i-l+1)
	}
	return res
}
