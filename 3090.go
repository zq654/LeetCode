package LeetCode

func maximumLengthSubstring(s string) int {
	//用Map记录字符出现次数
	occupyMap := make(map[byte]int)
	res := 0
	l := 0
	for i := 0; i < len(s); i++ {
		occupyMap[s[i]]++
		for occupyMap[s[i]] > 2 {
			occupyMap[s[l]]--
			l++
		}
		res = max(res, i-l+1)
	}
	return res
}
