package LeetCode

func countPalindromicSubsequence(s string) int {
	//维护前缀和后缀字符的数量
	//preArr 代表当index左侧的字符数量，sufArr 代表当index右侧的字符数量
	preArr := make([]int, 26)
	sufArr := make([]int, 26)
	for _, b := range s {
		sufArr[b-'a']++
	}
	res := 0
	for _, b := range s {
		preArr[b-'a']++
		sufArr[b-'a']--

	}
	return res
}
