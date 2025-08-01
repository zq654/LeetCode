package LeetCode

func lengthOfLongestSubstring(s string) int {
	//特殊值直接返回
	if s == "" {
		return 0
	}
	// 用res记录最大值
	res := 0
	// 左指针记录子串最左边的index
	left := 0
	//用map记录字母第一次出现的位置
	indexs := make(map[byte]int, len(s))
	for right := 0; right < len(s); right++ {
		char := s[right]
		index, ok := indexs[char]
		if !ok {
			indexs[char] = right
			res = max(res, right-left+1)
			continue
		}
		//若已经出现过 则计算长度 并将左index移到不与right值重复的index上
		indexs[char] = right
		//确保left不往回走
		left = max(index+1, left)
		length := right - left + 1
		res = max(length, res)
	}
	return res
}
