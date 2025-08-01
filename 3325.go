package LeetCode

func numberOfSubstrings(s string, k int) int {
	dif := 0 //记录窗口内有几个字符的数量大于k
	byteMap := make([]int, 26)
	res := 0
	l := 0
	for r := 0; r < len(s); r++ {
		byteMap[s[r]-'a']++
		if byteMap[s[r]-'a'] == k {
			dif++
		}
		for dif >= 1 {
			byteMap[s[l]-'a']--
			if byteMap[s[l]-'a'] == k-1 {
				dif--
			}
			l++
		}
		res += l
	}
	return res
}
